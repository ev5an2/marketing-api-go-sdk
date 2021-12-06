/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 商品信息
type ProductItemSpec struct {
	ProductOuterId            *string               `json:"product_outer_id,omitempty"`
	ProductName               *string               `json:"product_name,omitempty"`
	Description               *string               `json:"description,omitempty"`
	CreatedTime               *string               `json:"created_time,omitempty"`
	LastModifiedTime          *string               `json:"last_modified_time,omitempty"`
	ExpirationTime            *string               `json:"expiration_time,omitempty"`
	ImageUrl                  *string               `json:"image_url,omitempty"`
	AdditionalImageUrl        *[]string             `json:"additional_image_url,omitempty"`
	VideoDuration             *string               `json:"video_duration,omitempty"`
	VideoUrl                  *string               `json:"video_url,omitempty"`
	PlayCount                 *int64                `json:"play_count,omitempty"`
	PublishTime               *string               `json:"publish_time,omitempty"`
	AdditionalVideoUrl        *[]string             `json:"additional_video_url,omitempty"`
	PcPageUrl                 *string               `json:"pc_page_url,omitempty"`
	MobileH5PageUrl           *string               `json:"mobile_h5_page_url,omitempty"`
	AndroidPageUrl            *string               `json:"android_page_url,omitempty"`
	IosPageUrl                *string               `json:"ios_page_url,omitempty"`
	WechatPageUrl             *string               `json:"wechat_page_url,omitempty"`
	AdditionalMobileH5PageUrl *string               `json:"additional_mobile_h5_page_url,omitempty"`
	AdditionalAndroidPageUrl  *string               `json:"additional_android_page_url,omitempty"`
	AdditionalIosPageUrl      *string               `json:"additional_ios_page_url,omitempty"`
	AdditionalWechatPageUrl   *string               `json:"additional_wechat_page_url,omitempty"`
	MiniProgramId             *string               `json:"mini_program_id,omitempty"`
	MiniProgramPath           *string               `json:"mini_program_path,omitempty"`
	AdditionalMiniProgramId   *string               `json:"additional_mini_program_id,omitempty"`
	AdditionalMiniProgramPath *string               `json:"additional_mini_program_path,omitempty"`
	UniversalLink             *string               `json:"universal_link,omitempty"`
	AdditionalUniversalLink   *string               `json:"additional_universal_link,omitempty"`
	ProductShortName          *string               `json:"product_short_name,omitempty"`
	Price                     *float64              `json:"price,omitempty"`
	OriginalPrice             *float64              `json:"original_price,omitempty"`
	Discount                  *float64              `json:"discount,omitempty"`
	SalePrice                 *float64              `json:"sale_price,omitempty"`
	StartTime                 *string               `json:"start_time,omitempty"`
	EndTime                   *string               `json:"end_time,omitempty"`
	SalesVolume               *int64                `json:"sales_volume,omitempty"`
	StockVolume               *int64                `json:"stock_volume,omitempty"`
	Slogan                    *string               `json:"slogan,omitempty"`
	CustomLabel               *[]string             `json:"custom_label,omitempty"`
	FirstCategoryId           *int64                `json:"first_category_id,omitempty"`
	SecondCategoryId          *int64                `json:"second_category_id,omitempty"`
	ThirdCategoryId           *int64                `json:"third_category_id,omitempty"`
	FourthCategoryId          *int64                `json:"fourth_category_id,omitempty"`
	FirstCategoryName         *string               `json:"first_category_name,omitempty"`
	SecondCategoryName        *string               `json:"second_category_name,omitempty"`
	ThirdCategoryName         *string               `json:"third_category_name,omitempty"`
	FourthCategoryName        *string               `json:"fourth_category_name,omitempty"`
	FirstCategoryUrl          *string               `json:"first_category_url,omitempty"`
	SecondCategoryUrl         *string               `json:"second_category_url,omitempty"`
	ThirdCategoryUrl          *string               `json:"third_category_url,omitempty"`
	FourthCategoryUrl         *string               `json:"fourth_category_url,omitempty"`
	BrandId                   *int64                `json:"brand_id,omitempty"`
	ProductBrand              *string               `json:"product_brand,omitempty"`
	BrandUrl                  *string               `json:"brand_url,omitempty"`
	PromotionId               *int64                `json:"promotion_id,omitempty"`
	PromotionName             *string               `json:"promotion_name,omitempty"`
	PromotionUrl              *string               `json:"promotion_url,omitempty"`
	ShopId                    *int64                `json:"shop_id,omitempty"`
	ShopName                  *string               `json:"shop_name,omitempty"`
	ShopUrl                   *string               `json:"shop_url,omitempty"`
	ShopCustomInfo            *string               `json:"shop_custom_info,omitempty"`
	ShopIdList                *[]string             `json:"shop_id_list,omitempty"`
	ProductViewCount          *int64                `json:"product_view_count,omitempty"`
	FavoriteCount             *int64                `json:"favorite_count,omitempty"`
	Rating                    *float64              `json:"rating,omitempty"`
	FavourableCommentRate     *float64              `json:"favourable_comment_rate,omitempty"`
	ProductOwnerType          ProductOwnerType      `json:"product_owner_type,omitempty"`
	Author                    *string               `json:"author,omitempty"`
	FullText                  *string               `json:"full_text,omitempty"`
	LikeCount                 *int64                `json:"like_count,omitempty"`
	ForwardCount              *int64                `json:"forward_count,omitempty"`
	CommentCount              *int64                `json:"comment_count,omitempty"`
	AuthorFansCount           *int64                `json:"author_fans_count,omitempty"`
	SemanticLabels            *[]string             `json:"semantic_labels,omitempty"`
	DetailImg                 *[]string             `json:"detail_img,omitempty"`
	ShowCount                 *int64                `json:"show_count,omitempty"`
	PlayRate                  *float64              `json:"play_rate,omitempty"`
	MakeMoneyOnline           *string               `json:"make_money_online,omitempty"`
	LiveBroadcast             *string               `json:"live_broadcast,omitempty"`
	PlatformAttribute         *string               `json:"platform_attribute,omitempty"`
	CustomData                *string               `json:"custom_data,omitempty"`
	ShootingPicCount          *int64                `json:"shooting_pic_count,omitempty"`
	FineeditCount             *int64                `json:"fineedit_count,omitempty"`
	AlbumPhotoCount           *int64                `json:"album_photo_count,omitempty"`
	AlbumCount                *int64                `json:"album_count,omitempty"`
	FrameCount                *int64                `json:"frame_count,omitempty"`
	Country                   *string               `json:"country,omitempty"`
	Province                  *string               `json:"province,omitempty"`
	City                      *string               `json:"city,omitempty"`
	District                  *string               `json:"district,omitempty"`
	Address                   *[]string             `json:"address,omitempty"`
	ShootingSceneInCount      ShootingSceneInCount  `json:"shooting_scene_in_count,omitempty"`
	ShootingSceneOutCount     ShootingSceneOutCount `json:"shooting_scene_out_count,omitempty"`
	BrideClothing             BrideClothing         `json:"bride_clothing,omitempty"`
	GroomClothing             GroomClothing         `json:"groom_clothing,omitempty"`
	ClothCount                ClothCount            `json:"cloth_count,omitempty"`
	ShootingScene             ShootingScene         `json:"shooting_scene,omitempty"`
	SetSpecial                *[]string             `json:"set_special,omitempty"`
	ShootingStyleCom          *[]string             `json:"shooting_style_com,omitempty"`
	ShootingSceneDetail       *[]string             `json:"shooting_scene_detail,omitempty"`
	TourPlaceDomestic         TourPlaceDomestic     `json:"tour_place_domestic,omitempty"`
	TourPlaceForeign          TourPlaceForeign      `json:"tour_place_foreign,omitempty"`
	MaxInsuranceQuota         *float64              `json:"max_insurance_quota,omitempty"`
	FirstMonthInsuranceFee    *float64              `json:"first_month_insurance_fee,omitempty"`
	MaxQuota                  *float64              `json:"max_quota,omitempty"`
	AnnualizedRateOfReturn    *float64              `json:"annualized_rate_of_return,omitempty"`
	RiskType                  RiskTypeStatus        `json:"risk_type,omitempty"`
	Artist                    *[]string             `json:"artist,omitempty"`
}
