import React, { ReactElement } from 'react';
import * as yup from 'yup';
import { Checkbox, Form, PageSection, TextInput } from '@patternfly/react-core';
import usePageState from 'Containers/Integrations/hooks/usePageState';
import FormMessage from 'Components/PatternFly/FormMessage';
import FormLabelGroup from 'Containers/Integrations/IntegrationForm/FormLabelGroup';
import FormSaveButton from 'Components/PatternFly/FormSaveButton';
import FormCancelButton from 'Components/PatternFly/FormCancelButton';
import FormTestButton from 'Components/PatternFly/FormTestButton';
import { CloudSourceIntegration } from 'services/CloudSourceService';
import merge from 'lodash/merge';
import IntegrationFormActions from '../IntegrationFormActions';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

export const validationSchema = yup.object().shape({
    cloudSource: yup.object().shape({
        name: yup.string().trim().required('Integration name is required'),
        type: yup.string().matches(/TYPE_OCM/),
        credentials: yup.object().shape({
            secret: yup
                .string()
                .test('secret-test', 'Token is required', (value, context: yup.TestContext) => {
                    const requireSecretField =
                        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                        // @ts-ignore
                        context?.from[2]?.value?.updateCredentials || false;

                    if (!requireSecretField) {
                        return true;
                    }

                    const trimmedValue = value?.trim();
                    return !!trimmedValue;
                }),
        }),
        ocm: yup.object().shape({
            endpoint: yup.string().trim().required('Endpoint is required'),
        }),
        skipTestIntegration: yup.bool(),
    }),
    updatePassword: yup.bool(),
});

export type CloudSourceIntegrationFormValues = {
    cloudSource: CloudSourceIntegration;
    updateCredentials: boolean;
};
export const defaultValues: CloudSourceIntegrationFormValues = {
    cloudSource: {
        id: '',
        name: '',
        type: 'TYPE_OCM',
        credentials: {
            secret: '',
        },
        skipTestIntegration: true,
        ocm: {
            endpoint: 'https://api.openshift.com',
        },
    },
    updateCredentials: true,
};

function OcmIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<CloudSourceIntegration>): ReactElement {
    const formInitialValues = { ...defaultValues, ...initialValues };
    if (initialValues) {
        formInitialValues.cloudSource = merge({}, formInitialValues.cloudSource, initialValues);
        formInitialValues.cloudSource.credentials.secret = '';
        formInitialValues.updateCredentials = false;
    }
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<CloudSourceIntegrationFormValues>({
        initialValues: formInitialValues,
        validationSchema,
    });

    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    function onUpdateCredentialsChange(value, event) {
        setFieldValue('cloudSource.credentials.secret', '');
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormLabelGroup
                        isRequired
                        label="Integration name"
                        fieldId="cloudSource.name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="cloudSource.name"
                            value={values.cloudSource.name}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        isRequired
                        label="Endpoint"
                        fieldId="cloudSource.ocm.endpoint"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="cloudSource.ocm.endpoint"
                            name="cloudSource.ocm.endpoint"
                            value={values.cloudSource.ocm?.endpoint}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    {!isCreating && isEditable && (
                        <FormLabelGroup
                            fieldId="updateCredentials"
                            helperText="Enable this option to replace currently stored credentials (if any)"
                            errors={errors}
                        >
                            <Checkbox
                                label="Update stored credentials"
                                id="updateCredentials"
                                isChecked={values.updateCredentials}
                                onChange={onUpdateCredentialsChange}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                    )}
                    <FormLabelGroup
                        isRequired={values.updateCredentials}
                        label="API token"
                        fieldId="cloudSource.credentials.secret"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired={values.updateCredentials}
                            type="password"
                            id={`cloudSource.credentials.secret`}
                            value={values.cloudSource.credentials.secret}
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || !values.updateCredentials}
                            placeholder={
                                values.updateCredentials
                                    ? ''
                                    : 'Currently-stored token will be used.'
                            }
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default OcmIntegrationForm;
