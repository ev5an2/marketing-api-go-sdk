/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MarketingTargetAssetsUpdateRequest struct {
	OrganizationId     *int64            `json:"organization_id,omitempty"`
	MarketingAssetId   *int64            `json:"marketing_asset_id,omitempty"`
	MarketingAssetName *string           `json:"marketing_asset_name,omitempty"`
	Properties         *[]PropertyStruct `json:"properties,omitempty"`
}
