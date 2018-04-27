import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {
    Line,
    BarChart,
    Bar,
    Cell,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer
} from 'recharts';
import { format, subDays } from 'date-fns';
import * as Icon from 'react-feather';
import Slider from 'react-slick';
import { CarouselNextArrow, CarouselPrevArrow } from 'Components/CarouselArrows';
import { connect } from 'react-redux';
import { createSelector, createStructuredSelector } from 'reselect';

import TwoLevelPieChart from 'Components/visuals/TwoLevelPieChart';
import CustomLineChart from 'Components/visuals/CustomLineChart';
import DashboardBenchmarks from 'Containers/Dashboard/DashboardBenchmarks';
import SeverityTile from 'Containers/Dashboard/SeverityTile';
import TopRiskyDeployments from 'Containers/Dashboard/TopRiskyDeployments';
import cloneDeep from 'lodash/cloneDeep';
import { severityLabels } from 'messages/common';
import { selectors } from 'reducers';

//  @TODO: Have one source of truth for severity colors
const severityColorMap = {
    CRITICAL_SEVERITY: 'hsl(7, 100%, 55%)',
    HIGH_SEVERITY: 'hsl(349, 100%, 78%)',
    MEDIUM_SEVERITY: 'hsl(20, 100%, 78%)',
    LOW_SEVERITY: 'hsl(42, 100%, 84%)'
};

const severityPropType = PropTypes.oneOf([
    'CRITICAL_SEVERITY',
    'HIGH_SEVERITY',
    'MEDIUM_SEVERITY',
    'LOW_SEVERITY'
]);

const slickSettings = {
    dots: false,
    nextArrow: <CarouselNextArrow />,
    prevArrow: <CarouselPrevArrow />
};

const groupedViolationsPropType = PropTypes.arrayOf(
    PropTypes.shape({
        counts: PropTypes.arrayOf(
            PropTypes.shape({
                count: PropTypes.string.isRequired,
                severity: severityPropType
            })
        ),
        group: PropTypes.string.isRequired
    })
);

class DashboardPage extends Component {
    static propTypes = {
        violatonsByPolicyCategory: groupedViolationsPropType.isRequired,
        violationsByCluster: groupedViolationsPropType.isRequired,
        alertsByTimeseries: PropTypes.arrayOf(PropTypes.shape()).isRequired,
        benchmarks: PropTypes.arrayOf(PropTypes.shape()).isRequired,
        clustersByName: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types,
        deployments: PropTypes.arrayOf(PropTypes.object).isRequired,
        history: PropTypes.shape({
            push: PropTypes.func.isRequired
        }).isRequired
    };

    makeBarClickHandler = (clusterName, severity) => () => {
        // if clusters are not loaded yet, at least we can redirect to unfiltered violations
        const clusterQuery = clusterName !== '' ? `cluster=${clusterName}` : '';
        this.props.history.push(`/main/violations?severity=${severity}&${clusterQuery}`);
    };

    formatTimeseriesData = clusterData => {
        if (!clusterData) return '';
        // set a baseline zero'd object for the past week
        const baselineData = {};
        const xAxisBuckets = [];
        for (let i = 6; i >= 0; i -= 1) {
            const key = format(subDays(new Date(), i), 'MMM DD');
            baselineData[key] = 0;
            xAxisBuckets.push(key);
        }
        // set severities in timeAlertMap to have this zero'd data
        const timeAlertMap = {};
        const timeAlertInitialMap = {}; // this is the number of initial alerts that have come before
        Object.keys(severityColorMap).forEach(severity => {
            timeAlertMap[severity] = cloneDeep(baselineData);
            timeAlertInitialMap[severity] = 0;
        });

        // populate actual data into timeAlertMap
        clusterData.severities.forEach(severityObj => {
            const { severity, events } = severityObj;
            events.forEach(alert => {
                const time = format(parseInt(alert.time, 10), 'MMM DD');
                const alerts = timeAlertMap[severity][time];
                if (alerts !== undefined) {
                    switch (alert.type) {
                        case 'CREATED':
                            timeAlertMap[severity][time] += 1;
                            break;
                        case 'REMOVED':
                            timeAlertMap[severity][time] -= 1;
                            break;
                        default:
                            break;
                    }
                } else {
                    timeAlertInitialMap[severity] += 1;
                }
            });
        });

        Object.keys(severityColorMap).forEach(severity => {
            let runningSum = timeAlertInitialMap[severity];
            Object.keys(baselineData).forEach(time => {
                const prevVal = timeAlertMap[severity][time];
                timeAlertMap[severity][time] += runningSum;
                runningSum += prevVal;
            });
        });

        // set data format for line chart
        const cluster = {};
        cluster.data = Object.keys(baselineData).map(time => ({
            time,
            low: timeAlertMap.LOW_SEVERITY[time],
            medium: timeAlertMap.MEDIUM_SEVERITY[time],
            high: timeAlertMap.HIGH_SEVERITY[time],
            critical: timeAlertMap.CRITICAL_SEVERITY[time]
        }));
        cluster.name = clusterData.cluster;

        return cluster;
    };

    renderAlertsByTimeseries = () => {
        if (!this.props.alertsByTimeseries) return '';

        return (
            <div className="p-0 h-full w-full">
                <Slider {...slickSettings}>
                    {this.props.alertsByTimeseries.map(cluster => {
                        const { data, name } = this.formatTimeseriesData(cluster);
                        return (
                            <div className="h-64" key={name}>
                                <CustomLineChart
                                    data={data}
                                    name={name}
                                    xAxisDataKey="time"
                                    yAxisDataKey=""
                                >
                                    <Line
                                        type="monotone"
                                        dataKey="low"
                                        stroke={severityColorMap.LOW_SEVERITY}
                                    />
                                    <Line
                                        type="monotone"
                                        dataKey="medium"
                                        stroke={severityColorMap.MEDIUM_SEVERITY}
                                    />
                                    <Line
                                        type="monotone"
                                        dataKey="high"
                                        stroke={severityColorMap.HIGH_SEVERITY}
                                    />
                                    <Line
                                        type="monotone"
                                        dataKey="critical"
                                        stroke={severityColorMap.CRITICAL_SEVERITY}
                                    />
                                </CustomLineChart>
                            </div>
                        );
                    })}
                </Slider>
            </div>
        );
    };

    renderViolationsByCluster = () => {
        if (!this.props.violationsByCluster || !this.props.violationsByCluster.length) {
            return (
                <div className="flex flex-1 items-center justify-center">No Clusters Available</div>
            );
        }
        const clusterCharts = [];

        let i = 0;
        const limit = 4;
        while (i < this.props.violationsByCluster.length) {
            let j = i;
            let groupIndex = 0;
            const barCharts = [];
            while (j < this.props.violationsByCluster.length && groupIndex < limit) {
                const cluster = this.props.violationsByCluster[j];
                const dataPoint = {
                    name: cluster.group,
                    Critical: 0,
                    High: 0,
                    Medium: 0,
                    Low: 0
                };
                cluster.counts.forEach(d => {
                    dataPoint[severityLabels[d.severity]] = parseInt(d.count, 10);
                });
                barCharts.push(dataPoint);
                j += 1;
                groupIndex += 1;
            }
            clusterCharts.push(barCharts);
            i += 4;
        }
        return (
            <div className="p-0 h-full w-full">
                <Slider {...slickSettings}>
                    {clusterCharts.map((data, index) => (
                        <div key={index}>
                            <ResponsiveContainer className="flex-1 h-full w-full">
                                <BarChart
                                    data={data}
                                    margin={{
                                        top: 5,
                                        right: 30,
                                        left: 20,
                                        bottom: 5
                                    }}
                                >
                                    <XAxis dataKey="name" />
                                    <YAxis
                                        domain={[0, 'dataMax']}
                                        allowDecimals={false}
                                        label={{
                                            value: 'Count',
                                            angle: -90,
                                            position: 'insideLeft',
                                            textAnchor: 'middle'
                                        }}
                                    />
                                    <CartesianGrid strokeDasharray="3 3" />
                                    <Tooltip />
                                    <Legend
                                        horizontalAlign="right"
                                        wrapperStyle={{ lineHeight: '40px' }}
                                    />
                                    {Object.keys(severityLabels).map(severity => {
                                        const arr = [];
                                        const bar = (
                                            <Bar
                                                name={severityLabels[severity]}
                                                key={severityLabels[severity]}
                                                dataKey={severityLabels[severity]}
                                                fill={severityColorMap[severity]}
                                            >
                                                {data.map(entry => (
                                                    <Cell
                                                        key={entry.name}
                                                        className="cursor-pointer"
                                                        onClick={this.makeBarClickHandler(
                                                            entry.name,
                                                            severity
                                                        )}
                                                    />
                                                ))}
                                            </Bar>
                                        );
                                        arr.push(bar);
                                        return arr;
                                    })}
                                </BarChart>
                            </ResponsiveContainer>
                        </div>
                    ))}
                </Slider>
            </div>
        );
    };

    renderViolationsByPolicyCategory = () => {
        if (!this.props.violatonsByPolicyCategory) return '';
        return this.props.violatonsByPolicyCategory.map(policyType => {
            const data = policyType.counts.map(d => ({
                name: severityLabels[d.severity],
                value: parseInt(d.count, 10),
                color: severityColorMap[d.severity],
                onClick: () => {
                    this.props.history.push(
                        `/main/violations?category=${policyType.group}&severity=${d.severity}`
                    );
                }
            }));
            return (
                <div className="p-6 w-full lg:w-1/2" key={policyType.group}>
                    <div className="flex flex-col p-4 bg-white rounded-sm shadow">
                        <h2 className="flex items-center text-lg text-base font-sans text-base-600 py-4 tracking-wide">
                            <Icon.BarChart className="h-4 w-4 mr-3" />
                            {policyType.group}
                        </h2>
                        <div className="flex flex-1 m-4 h-64">
                            <TwoLevelPieChart data={data} />
                        </div>
                    </div>
                </div>
            );
        });
    };

    renderEnvironmentRisk = () => {
        const counts = {
            CRITICAL_SEVERITY: 0,
            HIGH_SEVERITY: 0,
            MEDIUM_SEVERITY: 0,
            LOW_SEVERITY: 0
        };
        this.props.violationsByCluster.forEach(cluster => {
            cluster.counts.forEach(d => {
                const count = parseInt(d.count, 10);
                counts[d.severity] += count;
            });
        });
        const severities = Object.keys(counts);
        return (
            <div className="flex flex-1 flex-col w-full">
                <h2 className="flex items-center text-xl text-base font-sans text-base-600 pb-8 tracking-wide font-500">
                    Environment Risk
                </h2>
                <div className="flex">
                    {severities.map((severity, i) => (
                        <SeverityTile
                            severity={severity}
                            count={counts[severity]}
                            color={severityColorMap[severity]}
                            index={i}
                            key={severity}
                        />
                    ))}
                </div>
            </div>
        );
    };

    renderBenchmarks = () => (
        <div className="p-0 h-full w-full dashboard-benchmarks">
            <Slider {...slickSettings}>
                {this.props.benchmarks.map((cluster, index) => (
                    <div key={index}>
                        <DashboardBenchmarks cluster={cluster} />
                    </div>
                ))}
            </Slider>
        </div>
    );

    renderTopRiskyDeployments = () => {
        if (!this.props.deployments) return '';
        const deployments = this.props.deployments.slice(0, 5);
        return <TopRiskyDeployments deployments={deployments} />;
    };

    render() {
        return (
            <section className="w-full h-full transition">
                <div className="flex bg-white border-b border-primary-500">
                    <div className="w-1/2 p-6">{this.renderEnvironmentRisk()}</div>
                    <div className="w-1/2 p-6 border-l border-primary-200">
                        {this.renderBenchmarks()}
                    </div>
                </div>
                <div className="overflow-auto bg-base-100">
                    <div className="flex flex-col w-full">
                        <div className="flex w-full flex-wrap">
                            <div className="p-6 md:w-full lg:w-1/2">
                                <div className="flex flex-col p-4 bg-white rounded-sm shadow h-full">
                                    <h2 className="flex items-center text-lg text-base font-sans text-base-600 py-4 tracking-wide">
                                        <Icon.Layers className="h-4 w-4 mr-3" />
                                        Violations by Cluster
                                    </h2>
                                    <div className="flex flex-1 m-4 h-64">
                                        {this.renderViolationsByCluster()}
                                    </div>
                                </div>
                            </div>
                            <div className="p-6 md:w-full lg:w-1/2">
                                <div className="flex flex-col p-4 bg-white rounded-sm shadow">
                                    <h2 className="flex items-center text-lg text-base font-sans text-base-600 py-4 tracking-wide">
                                        <Icon.AlertTriangle className="h-4 w-4 mr-3" />
                                        Active Violations by Time
                                    </h2>
                                    <div className="flex flex-1 m-4 h-64">
                                        {this.renderAlertsByTimeseries()}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-col w-full">
                        <div className="flex w-full flex-wrap">
                            {this.renderViolationsByPolicyCategory()}
                            <div className="p-6 md:w-full lg:w-1/2">
                                {this.renderTopRiskyDeployments()}
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        );
    }
}

const getClustersByName = createSelector([selectors.getClusters], clusters =>
    clusters.reduce(
        (result, cluster) => ({
            ...result,
            [cluster.name]: cluster
        }),
        {}
    )
);

const mapStateToProps = createStructuredSelector({
    violatonsByPolicyCategory: selectors.getAlertCountsByPolicyCategories,
    violationsByCluster: selectors.getAlertCountsByCluster,
    alertsByTimeseries: selectors.getAlertsByTimeseries,
    benchmarks: selectors.getBenchmarksByCluster,
    deployments: selectors.getDeployments,
    clustersByName: getClustersByName
});

export default connect(mapStateToProps)(DashboardPage);
