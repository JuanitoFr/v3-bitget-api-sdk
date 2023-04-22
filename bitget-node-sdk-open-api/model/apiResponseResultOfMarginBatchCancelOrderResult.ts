/**
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { MarginBatchCancelOrderResult } from './marginBatchCancelOrderResult';

export class ApiResponseResultOfMarginBatchCancelOrderResult {
    /**
    * code
    */
    'code'?: string;
    'data'?: MarginBatchCancelOrderResult;
    /**
    * msg
    */
    'msg'?: string;
    /**
    * requestTime
    */
    'requestTime'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "code",
            "baseName": "code",
            "type": "string"
        },
        {
            "name": "data",
            "baseName": "data",
            "type": "MarginBatchCancelOrderResult"
        },
        {
            "name": "msg",
            "baseName": "msg",
            "type": "string"
        },
        {
            "name": "requestTime",
            "baseName": "requestTime",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return ApiResponseResultOfMarginBatchCancelOrderResult.attributeTypeMap;
    }
}
