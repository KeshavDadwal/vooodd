package api

import (
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Querier
	//tokenMaker token.Maker
	router *mux.Router
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Querier, config util.Config) (*Server, error) {

	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()

	//POST
	router.HandleFunc("/users", server.createUser).Methods("POST")
	router.HandleFunc("/country", server.createCountry).Methods("POST")
	//router.HandleFunc("/users/login", server.loginUser).Methods("POST")
	router.HandleFunc("/currencies",server.createCurrency).Methods("POST")
	router.HandleFunc("/language",server.createLanguage).Methods("POST")
	router.HandleFunc("/role",server.createRole).Methods("POST")
	router.HandleFunc("/userrole",server.createUserRole).Methods("POST")
	router.HandleFunc("/brands",server.createBrand).Methods("POST")
	router.HandleFunc("/brandlanguage",server.createBrandLanguage).Methods("POST")
	router.HandleFunc("/competitor",server.createCompetitor).Methods("POST")
	router.HandleFunc("/competitorlanguage",server.createCompetitorLanguage).Methods("POST")
	router.HandleFunc("/contact",server.createContact).Methods("POST")
	router.HandleFunc("/contentmanagement",server.createContentManagement).Methods("POST")
	router.HandleFunc("/contentmanagementbanner",server.createContentManagementBanner).Methods("POST")
	router.HandleFunc("/contentmanagementbannerlanguage",server.createContentManagementBannerLanguage).Methods("POST")
	router.HandleFunc("/contentmanagementlanguage",server.createContentManagementLanguage).Methods("POST")
	router.HandleFunc("/contentmanagementmobile",server.createContentManagementMobile).Methods("POST")
	router.HandleFunc("/contentmanagementmobilelanguage",server.createContentManagementMobileLanguage).Methods("POST")
	router.HandleFunc("/store",server.createStore).Methods("POST")
	router.HandleFunc("/unit",server.createUnit).Methods("POST")
	router.HandleFunc("/productcategory",server.createProductCategory).Methods("POST")
	router.HandleFunc("/product",server.createProduct).Methods("POST")
	router.HandleFunc("/coupon",server.createCoupon).Methods("POST")
	router.HandleFunc("/couponproduct",server.createCouponProduct).Methods("POST")
	router.HandleFunc("/couponlanguage",server.createCouponLanguage).Methods("POST")
	router.HandleFunc("/feedback",server.createFeedback).Methods("POST")
	router.HandleFunc("/deal",server.createDeal).Methods("POST")
	router.HandleFunc("/deallanguage",server.createDealsLanguage).Methods("POST")
	router.HandleFunc("/homepageimage",server.createHomePageImage).Methods("POST")
	router.HandleFunc("/homepageimagelanguage",server.createHomePageImageLanguage).Methods("POST")
	router.HandleFunc("/homepageicon",server.createHomePageIcon).Methods("POST")
	router.HandleFunc("/homepageiconlanguage",server.createHomePageIconLanguage).Methods("POST")
	router.HandleFunc("/homebanner",server.createHomeBanner).Methods("POST")
	router.HandleFunc("/offer",server.createOffer).Methods("POST")
	router.HandleFunc("/offercategory",server.createOfferCategory).Methods("POST")
	router.HandleFunc("/producttest",server.createProductTest).Methods("POST")
	router.HandleFunc("/productstorelocation",server.createProductStoreLocation).Methods("POST")
	router.HandleFunc("/productlanguage",server.createProductLanguage).Methods("POST")
	router.HandleFunc("/productimage",server.createProductImage).Methods("POST")
	router.HandleFunc("/productcategorymap",server.createProductCategoryMap).Methods("POST")
	router.HandleFunc("/productcategorylanguage",server.createProductCategoryLanguage).Methods("POST")
	router.HandleFunc("/sharecount",server.createShareCount).Methods("POST")
	router.HandleFunc("/storenamecategory",server.CreateStoreNameCategory).Methods("POST")
	router.HandleFunc("/storelangauge",server.createStoreLanguage).Methods("POST")
	router.HandleFunc("/storecategory",server.createStoreCategory).Methods("POST")
	router.HandleFunc("/unitlanguage",server.createUnitLanguage).Methods("POST")
	router.HandleFunc("/usershoppinglistname",server.createUserShoppingListName).Methods("POST")
	router.HandleFunc("/userfavouritestore",server.createUserFavoriteStore).Methods("POST")
	router.HandleFunc("/userevent",server.createUserEvent).Methods("POST")
	router.HandleFunc("/userdevice",server.createUserDevice).Methods("POST")
	router.HandleFunc("/usercoupon",server.createUserCoupon).Methods("POST")
	router.HandleFunc("/zone",server.createZone).Methods("POST")
	router.HandleFunc("/mall",server.createMall).Methods("POST")
	router.HandleFunc("/malllangauge",server.createMallLanguage).Methods("POST")
	router.HandleFunc("/mallview",server.createMallViews).Methods("POST")
	router.HandleFunc("/storelocation",server.createStoreLocation).Methods("POST")
	router.HandleFunc("/storelocationoffer",server.createStoreLocationOffers).Methods("POST")
	router.HandleFunc("/storelocationholiday",server.createStoreLocationHolidays).Methods("POST")
	router.HandleFunc("/storelocationcoupon",server.createStoreLocationCoupon).Methods("POST")
	router.HandleFunc("/storebusinesshour",server.createStoreBusinessHours).Methods("POST")
	router.HandleFunc("/subscription",server.createSubscription).Methods("POST")
	router.HandleFunc("/couponcondition",server.createCouponCondition).Methods("POST")
	router.HandleFunc("/dislikecount",server.createDislikeCount).Methods("POST")
	router.HandleFunc("/importflow",server.createImportFlow).Methods("POST")
	router.HandleFunc("/likecount",server.createLikeCount).Methods("POST")
	router.HandleFunc("/offerview",server.createOfferView).Methods("POST")
	router.HandleFunc("/pushnotification",server.createPushNotification).Methods("POST")
	router.HandleFunc("/userwishlist",server.createUserWishlistItem).Methods("POST")
	router.HandleFunc("/userstore",server.createUserStore).Methods("POST")
	router.HandleFunc("/usershoppinglist",server.createUserShopingList).Methods("POST")
	router.HandleFunc("/usersetting",server.createUserSetting).Methods("POST")
	router.HandleFunc("/usermalls",server.createUserMall).Methods("POST")
	router.HandleFunc("/usermallpermission",server.createUserMallPermission).Methods("POST")
	router.HandleFunc("/userlog",server.createUserLog).Methods("POST")
	router.HandleFunc("/useraction",server.createUserAction).Methods("POST")

	// Get

	//router.HandleFunc("brandlanguage/{id}",server.getBrandLanguagesById).Methods("GET")
	router.HandleFunc("/brandlanguage/{id}", server.getBrandLanguagesById).Methods("GET")
	router.HandleFunc("/brandlanguage",server.getAllBrandLanguages).Methods("GET")
	router.HandleFunc("/brand/{id}", server.getBrandById).Methods("GET")
	router.HandleFunc("/brand",server.getAllBrand).Methods("GET")
	router.HandleFunc("/competitorlanguage/{id}", server.getCompetitorLanguageById).Methods("GET")
	router.HandleFunc("/competitorlanguage", server.getAllCompetitorLanguage).Methods("GET")
	router.HandleFunc("/competitor/{id}", server.getCompetitorsById).Methods("GET")
	router.HandleFunc("/competitor", server.getAllCompetitors).Methods("GET")
	router.HandleFunc("/contact/{id}", server.getContactById).Methods("GET")
	router.HandleFunc("/contact", server.getAllContact).Methods("GET")
	router.HandleFunc("/contentmanagementbannerlanguage/{id}", server.getContentManagementBannerLanguageById).Methods("GET")
	router.HandleFunc("/contentmanagementbannerlanguage", server.getAllContentManagementBannerLanguage).Methods("GET")
	router.HandleFunc("/contentmanagementbanner/{id}", server.getContentManagementBannerById).Methods("GET")
	router.HandleFunc("/contentmanagementbanner", server.getAllContentManagementBanner).Methods("GET")
	router.HandleFunc("/contentmanagementlanguage/{id}", server.getContentManagementLanguageById).Methods("GET")
	router.HandleFunc("/contentmanagementlanguage", server.getAllContentManagementLanguage).Methods("GET")
	router.HandleFunc("/contentmanagementmobilelanguage/{id}", server.getContentManagementMobileLanguage).Methods("GET")
	router.HandleFunc("/contentmanagementmobilelanguage", server.getAllContentManagementMobileLanguage).Methods("GET")
	router.HandleFunc("/contentmanagementmobile/{id}", server.getContentManagementMobileById).Methods("GET")
	router.HandleFunc("/contentmanagementmobile", server.getAllContentManagementMobile).Methods("GET")
	router.HandleFunc("/contentmanagement/{id}", server.getContentManagementById).Methods("GET")
	router.HandleFunc("/contentmanagement", server.getAllContentManagement).Methods("GET")
	router.HandleFunc("/country/{id}", server.getCountryById).Methods("GET")
	router.HandleFunc("/country", server.getAllCountry).Methods("GET")
	router.HandleFunc("/couponcondition/{id}", server.getCouponConditionById).Methods("GET")
	router.HandleFunc("/couponcondition", server.getAllCouponCondition).Methods("GET")
	router.HandleFunc("/couponlanguage", server.getAllCouponLanguage).Methods("GET")
	router.HandleFunc("/couponproduct", server.getAllCouponProduct).Methods("GET")
	router.HandleFunc("/coupon/{id}", server.getCouponById).Methods("GET")
	router.HandleFunc("/coupon", server.getAllCoupon).Methods("GET")
	router.HandleFunc("/currency/{id}", server.getCurrencyById).Methods("GET")
	router.HandleFunc("/currency", server.getAllCurrency).Methods("GET")
	router.HandleFunc("/deallanguage/{id}", server.getDealsLanguageById).Methods("GET")
	router.HandleFunc("/deallanguage", server.getAllDealsLanguage).Methods("GET")
	router.HandleFunc("/deal/{id}", server.getDealById).Methods("GET")
	router.HandleFunc("/deal", server.getAllDeals).Methods("GET")
	router.HandleFunc("/dislikecount", server.getAllDislikeCount).Methods("GET")
	router.HandleFunc("/feedback", server.getAllFeedback).Methods("GET")
	router.HandleFunc("/homebanner/{id}", server.getHomeBannerById).Methods("GET")
	router.HandleFunc("/homebanner", server.getAllHomeBanner).Methods("GET")
	router.HandleFunc("/homepageicon/{id}", server.getHomePageIconById).Methods("GET")
	router.HandleFunc("/homepageicon", server.getAllHomePageIcon).Methods("GET")
	router.HandleFunc("/homepageiconlanguage", server.getAllHomePageIconLanguage).Methods("GET")
	router.HandleFunc("/homepageimagelanguage", server.getAllHomePageImageLanguage).Methods("GET")
	router.HandleFunc("/homepageimage", server.getAllHomePageImage).Methods("GET")
	router.HandleFunc("/importflow", server.getAllImportFlow).Methods("GET")
	router.HandleFunc("/language", server.getAllLanguage).Methods("GET")
	router.HandleFunc("/likecount", server.getAllLikeCount).Methods("GET")
//	router.HandleFunc("/tokens/renew_access", server.refreshToken)

	//router.HandleFunc("/upload/video", server.uploadVideoToS3).Methods("POST")
	//router.HandleFunc("/videos", server.listAllVideos).Methods("GET")
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	err := http.ListenAndServe(address, server.router)
	if err != nil {
		return err
	}
	return nil
}
