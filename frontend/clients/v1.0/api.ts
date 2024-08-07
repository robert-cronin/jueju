/* tslint:disable */
/* eslint-disable */
/**
 * JueJu API
 * This is the JueJu API
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * Error code
     * @type {number}
     * @memberof ModelError
     */
    'code': number;
    /**
     * Error message
     * @type {string}
     * @memberof ModelError
     */
    'message': string;
}
/**
 * 
 * @export
 * @interface PoemRequest
 */
export interface PoemRequest {
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'user_id': string;
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'prompt': string;
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'poem'?: string;
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'status': PoemRequestStatusEnum;
    /**
     * Number of attempts made for this poem request
     * @type {number}
     * @memberof PoemRequest
     */
    'attempt_count'?: number;
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'created_at': string;
    /**
     * 
     * @type {string}
     * @memberof PoemRequest
     */
    'updated_at': string;
}

export const PoemRequestStatusEnum = {
    Pending: 'pending',
    Completed: 'completed',
    Failed: 'failed'
} as const;

export type PoemRequestStatusEnum = typeof PoemRequestStatusEnum[keyof typeof PoemRequestStatusEnum];

/**
 * 
 * @export
 * @interface PoemRequestInput
 */
export interface PoemRequestInput {
    /**
     * 
     * @type {string}
     * @memberof PoemRequestInput
     */
    'prompt': string;
}
/**
 * 
 * @export
 * @interface RequestPoem403Response
 */
export interface RequestPoem403Response {
    /**
     * 
     * @type {string}
     * @memberof RequestPoem403Response
     */
    'error'?: string;
    /**
     * 
     * @type {number}
     * @memberof RequestPoem403Response
     */
    'credits_required'?: number;
    /**
     * 
     * @type {number}
     * @memberof RequestPoem403Response
     */
    'credits_available'?: number;
}
/**
 * 
 * @export
 * @interface User
 */
export interface User {
    /**
     * User ID
     * @type {string}
     * @memberof User
     */
    'id': string;
    /**
     * Auth0 ID
     * @type {string}
     * @memberof User
     */
    'auth0_id': string;
    /**
     * User email
     * @type {string}
     * @memberof User
     */
    'email': string;
    /**
     * Email verification status
     * @type {boolean}
     * @memberof User
     */
    'email_verified'?: boolean;
    /**
     * User name
     * @type {string}
     * @memberof User
     */
    'name'?: string;
    /**
     * User nickname
     * @type {string}
     * @memberof User
     */
    'nickname'?: string;
    /**
     * User avatar
     * @type {string}
     * @memberof User
     */
    'picture'?: string;
    /**
     * Account creation timestamp
     * @type {string}
     * @memberof User
     */
    'created_at'?: string;
    /**
     * Last update timestamp
     * @type {string}
     * @memberof User
     */
    'updated_at'?: string;
    /**
     * Last login timestamp
     * @type {string}
     * @memberof User
     */
    'last_login'?: string;
    /**
     * Number of poem credits available to the user
     * @type {number}
     * @memberof User
     */
    'poem_credits'?: number;
    /**
     * Timestamp of the last credit reset
     * @type {string}
     * @memberof User
     */
    'last_credit_reset'?: string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Callback from Auth0
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        callback: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/auth/callback`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get user information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getUser: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/user`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication cookieAuth required


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get user\'s poem requests
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getUserPoemRequests: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/poems`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication cookieAuth required


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Logs user into the system via Auth0
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/auth/login`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Logs out current logged in user session
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        logout: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/logout`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication cookieAuth required


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Request a new poem
         * @param {PoemRequestInput} poemRequestInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestPoem: async (poemRequestInput: PoemRequestInput, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'poemRequestInput' is not null or undefined
            assertParamExists('requestPoem', 'poemRequestInput', poemRequestInput)
            const localVarPath = `/poems`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication cookieAuth required


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(poemRequestInput, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Callback from Auth0
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async callback(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Error>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.callback(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.callback']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Get user information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getUser(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<User>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getUser(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.getUser']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Get user\'s poem requests
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getUserPoemRequests(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<PoemRequest>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getUserPoemRequests(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.getUserPoemRequests']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Logs user into the system via Auth0
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async login(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Error>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.login(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.login']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Logs out current logged in user session
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async logout(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Error>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.logout(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.logout']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Request a new poem
         * @param {PoemRequestInput} poemRequestInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async requestPoem(poemRequestInput: PoemRequestInput, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<PoemRequest>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.requestPoem(poemRequestInput, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['DefaultApi.requestPoem']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @summary Callback from Auth0
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        callback(options?: any): AxiosPromise<Error> {
            return localVarFp.callback(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get user information
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getUser(options?: any): AxiosPromise<User> {
            return localVarFp.getUser(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get user\'s poem requests
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getUserPoemRequests(options?: any): AxiosPromise<Array<PoemRequest>> {
            return localVarFp.getUserPoemRequests(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Logs user into the system via Auth0
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(options?: any): AxiosPromise<Error> {
            return localVarFp.login(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Logs out current logged in user session
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        logout(options?: any): AxiosPromise<Error> {
            return localVarFp.logout(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Request a new poem
         * @param {PoemRequestInput} poemRequestInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestPoem(poemRequestInput: PoemRequestInput, options?: any): AxiosPromise<PoemRequest> {
            return localVarFp.requestPoem(poemRequestInput, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @summary Callback from Auth0
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public callback(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).callback(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get user information
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getUser(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getUser(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get user\'s poem requests
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getUserPoemRequests(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getUserPoemRequests(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Logs user into the system via Auth0
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public login(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).login(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Logs out current logged in user session
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public logout(options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).logout(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Request a new poem
     * @param {PoemRequestInput} poemRequestInput 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public requestPoem(poemRequestInput: PoemRequestInput, options?: RawAxiosRequestConfig) {
        return DefaultApiFp(this.configuration).requestPoem(poemRequestInput, options).then((request) => request(this.axios, this.basePath));
    }
}



