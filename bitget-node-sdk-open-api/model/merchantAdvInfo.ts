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
import { FiatPaymentInfo } from './fiatPaymentInfo';
import { MerchantAdvUserLimitInfo } from './merchantAdvUserLimitInfo';

export class MerchantAdvInfo {
    'advId'?: string;
    'advNo'?: string;
    'amount'?: string;
    'coin'?: string;
    'coinPrecision'?: string;
    'ctime'?: string;
    'dealAmount'?: string;
    'fiatCode'?: string;
    'fiatPrecision'?: string;
    'fiatSymbol'?: string;
    'hide'?: string;
    'maxAmount'?: string;
    'minAmount'?: string;
    'payDuration'?: string;
    'paymentMethod'?: Array<FiatPaymentInfo>;
    'price'?: string;
    'remark'?: string;
    'status'?: string;
    'turnoverNum'?: string;
    'turnoverRate'?: string;
    'type'?: string;
    'userLimit'?: MerchantAdvUserLimitInfo;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "advId",
            "baseName": "advId",
            "type": "string"
        },
        {
            "name": "advNo",
            "baseName": "advNo",
            "type": "string"
        },
        {
            "name": "amount",
            "baseName": "amount",
            "type": "string"
        },
        {
            "name": "coin",
            "baseName": "coin",
            "type": "string"
        },
        {
            "name": "coinPrecision",
            "baseName": "coinPrecision",
            "type": "string"
        },
        {
            "name": "ctime",
            "baseName": "ctime",
            "type": "string"
        },
        {
            "name": "dealAmount",
            "baseName": "dealAmount",
            "type": "string"
        },
        {
            "name": "fiatCode",
            "baseName": "fiatCode",
            "type": "string"
        },
        {
            "name": "fiatPrecision",
            "baseName": "fiatPrecision",
            "type": "string"
        },
        {
            "name": "fiatSymbol",
            "baseName": "fiatSymbol",
            "type": "string"
        },
        {
            "name": "hide",
            "baseName": "hide",
            "type": "string"
        },
        {
            "name": "maxAmount",
            "baseName": "maxAmount",
            "type": "string"
        },
        {
            "name": "minAmount",
            "baseName": "minAmount",
            "type": "string"
        },
        {
            "name": "payDuration",
            "baseName": "payDuration",
            "type": "string"
        },
        {
            "name": "paymentMethod",
            "baseName": "paymentMethod",
            "type": "Array<FiatPaymentInfo>"
        },
        {
            "name": "price",
            "baseName": "price",
            "type": "string"
        },
        {
            "name": "remark",
            "baseName": "remark",
            "type": "string"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "string"
        },
        {
            "name": "turnoverNum",
            "baseName": "turnoverNum",
            "type": "string"
        },
        {
            "name": "turnoverRate",
            "baseName": "turnoverRate",
            "type": "string"
        },
        {
            "name": "type",
            "baseName": "type",
            "type": "string"
        },
        {
            "name": "userLimit",
            "baseName": "userLimit",
            "type": "MerchantAdvUserLimitInfo"
        }    ];

    static getAttributeTypeMap() {
        return MerchantAdvInfo.attributeTypeMap;
    }
}

