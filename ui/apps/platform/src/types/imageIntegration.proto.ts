import { ImageIntegrationType } from './integration';

export type BaseImageIntegration = {
    id: string;
    name: string;
    type: ImageIntegrationType;
    categories: ImageIntegrationCategory[];
    autogenerated: boolean;
    clusterId: string;
    skipTestIntegration: boolean;
};

export type ImageIntegrationCategory = 'REGISTRY' | 'SCANNER' | 'NODE_SCANNER';

// For strict type checking of frontend code, although not specified in proto.
export type CategoriesForClairifyScanner = 'SCANNER' | 'NODE_SCANNER';
export type CategoriesForRegistryScanner = 'REGISTRY' | 'SCANNER';

export type ImageIntegration =
    | ArtifactoryImageIntegration
    | ArtifactRegistryImageIntegration
    | AzureImageIntegration
    | ClairImageIntegration
    | ClairifyImageIntegration
    | DockerImageIntegration
    | EcrImageIntegration
    | GoogleImageIntegration
    | IbmImageIntegration
    | NexusImageIntegration
    | QuayImageIntegration
    | RhelImageIntegration;

export type ArtifactoryImageIntegration = {
    type: 'artifactory';
} & BaseImageIntegration;

export type ArtifactRegistryImageIntegration = {
    type: 'artifactregistry';
} & BaseImageIntegration;

export type AzureImageIntegration = {
    type: 'azure';
} & BaseImageIntegration;

export type ClairImageIntegration = {
    type: 'clair';
    clair: ClairConfig;
} & BaseImageIntegration;

export type ClairConfig = {
    endpoint: string;
    insecure: boolean;
};

export type ClairifyImageIntegration = {
    type: 'clairify';
    categories: CategoriesForClairifyScanner[];
    clairify: ClairifyConfig;
} & BaseImageIntegration;

export type ClairifyConfig = {
    endpoint: string;
    grpcEndpoint: string;
    numConcurrentScans: number; // int32
};

export type DockerImageIntegration = {
    type: 'docker';
    docker: DockerConfig;
} & BaseImageIntegration;

export type DockerConfig = {
    endpoint: string; // scrub: dependent
    username: string; // scrub: dependent
    // The password for the integration. The server will mask the value of this credential in responses and logs.
    password: string; // scrub: always
    insecure: boolean;
};

export type EcrImageIntegration = {
    type: 'ecr';
    ecr: EcrConfig;
} & BaseImageIntegration;

export type EcrConfig = {
    registryId: string;
    // The access key ID for the integration. The server will mask the value of this credential in responses and logs.
    accessKeyId: string; // scrub: always
    // The secret access key for the integration. The server will mask the value of this credential in responses and logs.
    secretAccessKey: string; // scrub: always
    region: string;
    useIam: boolean; // scrub: dependent
    endpoint: string; // scrub: dependent
    useAssumeRole: boolean;
    assumeRoleId: string;
    assumeRoleExternalId: string;
    authorizationData: EcrAuthorizationData;
};

export type EcrAuthorizationData = {
    username: string;
    passwordstring; // scrub: always
    expiresAt: string; // ISO 8601 date string
};

export type GoogleImageIntegration = {
    type: 'google';
    categories: CategoriesForRegistryScanner[];
    google: GoogleConfig;
} & BaseImageIntegration;

export type GoogleConfig = {
    endpoint: string; // scrub: dependent
    // The service account for the integration. The server will mask the value of this credential in responses and logs.
    serviceAccount: string; // scrub: always
    project: string;
};

export type IbmImageIntegration = {
    type: 'ibm';
    ibm: IbmRegistryConfig;
} & BaseImageIntegration;

export type IbmRegistryConfig = {
    endpoint: string;
    // The API key for the integration. The server will mask the value of this credential in responses and logs.
    apiKey: string;
};

export type NexusImageIntegration = {
    type: 'nexus';
} & BaseImageIntegration;

export type QuayImageIntegration = {
    type: 'quay';
    categories: CategoriesForRegistryScanner[];
    quay: QuayConfig;
} & BaseImageIntegration;

export type QuayConfig = {
    endpoint: string; // scrub:dependent
    // The OAuth token for the integration. The server will mask the value of this credential in responses and logs.
    oauthToken: string; // scrub: always
    insecure: boolean;
    // For registry integrations, Quay recommends using robot accounts. oauthToken will continue to be used for scanner integration.
    registryRobotCredentials: QuayRobotAccount | null;
};

// Robot account is Quay's named tokens that can be granted permissions on multiple repositories under an organization.
// It's Quay's recommended authentication model when possible (i.e. registry integration)
export type QuayRobotAccount = {
    username: string;
    // The server will mask the value of this password in responses and logs.
    password: string; // scrub: always
};

export type RhelImageIntegration = {
    type: 'rhel';
} & BaseImageIntegration;
