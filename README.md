# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BrandsApi* | [**ApiV2BrandsGet**](docs/BrandsApi.md#apiv2brandsget) | **Get** /api/v2/brands | 
*CountriesApi* | [**ApiV2BrandCountriesGet**](docs/CountriesApi.md#apiv2brandcountriesget) | **Get** /api/v2/{brand}/countries | 
*CountriesApi* | [**ApiV2CountriesGet**](docs/CountriesApi.md#apiv2countriesget) | **Get** /api/v2/countries | 
*OrdersApi* | [**ApiV2BrandOrdersCountBySourceYearMonthDayGet**](docs/OrdersApi.md#apiv2brandorderscountbysourceyearmonthdayget) | **Get** /api/v2/{brand}/orders/count-by-source/{year}/{month}/{day} | 
*OrdersApi* | [**ApiV2OrdersCountBySourceYearMonthDayGet**](docs/OrdersApi.md#apiv2orderscountbysourceyearmonthdayget) | **Get** /api/v2/orders/count-by-source/{year}/{month}/{day} | 
*ProductsApi* | [**ApiV2BrandCountryIdUnitIdProductsTopGet**](docs/ProductsApi.md#apiv2brandcountryidunitidproductstopget) | **Get** /api/v2/{brand}/{countryId}/{unitId}/products/top | 
*ProductsApi* | [**ApiV2BrandProductsTopGet**](docs/ProductsApi.md#apiv2brandproductstopget) | **Get** /api/v2/{brand}/products/top | 
*RatingsApi* | [**ApiV2BrandRatingsGet**](docs/RatingsApi.md#apiv2brandratingsget) | **Get** /api/v2/{brand}/ratings | 
*RatingsApi* | [**ApiV2RatingsGet**](docs/RatingsApi.md#apiv2ratingsget) | **Get** /api/v2/ratings | 
*RevenueApi* | [**ApiV2BrandCountryIdRevenueDailyGet**](docs/RevenueApi.md#apiv2brandcountryidrevenuedailyget) | **Get** /api/v2/{brand}/{countryId}/revenue/daily | 
*RevenueApi* | [**ApiV2BrandRevenueDailyYearMonthDayGet**](docs/RevenueApi.md#apiv2brandrevenuedailyyearmonthdayget) | **Get** /api/v2/{brand}/revenue/daily/{year}/{month}/{day} | 
*RevenueApi* | [**ApiV2BrandRevenueMonthsLastGet**](docs/RevenueApi.md#apiv2brandrevenuemonthslastget) | **Get** /api/v2/{brand}/revenue/months/last | 
*RevenueApi* | [**ApiV2BrandRevenueMonthsYearMonthGet**](docs/RevenueApi.md#apiv2brandrevenuemonthsyearmonthget) | **Get** /api/v2/{brand}/revenue/months/{year}/{month} | 
*RevenueApi* | [**ApiV2BrandRevenueWeeklyTopGet**](docs/RevenueApi.md#apiv2brandrevenueweeklytopget) | **Get** /api/v2/{brand}/revenue/weekly/top | 
*RevenueApi* | [**ApiV2RevenueDailyYearMonthDayGet**](docs/RevenueApi.md#apiv2revenuedailyyearmonthdayget) | **Get** /api/v2/revenue/daily/{year}/{month}/{day} | 
*RevenueApi* | [**ApiV2RevenueMonthsLastGet**](docs/RevenueApi.md#apiv2revenuemonthslastget) | **Get** /api/v2/revenue/months/last | 
*RevenueApi* | [**ApiV2RevenueMonthsYearMonthGet**](docs/RevenueApi.md#apiv2revenuemonthsyearmonthget) | **Get** /api/v2/revenue/months/{year}/{month} | 
*RevenueApi* | [**ApiV2RevenueWeeklyTopGet**](docs/RevenueApi.md#apiv2revenueweeklytopget) | **Get** /api/v2/revenue/weekly/top | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdDailyYearMonthDayGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitiddailyyearmonthdayget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/daily/{year}/{month}/{day} | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdHistoryGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidhistoryget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/history | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdHistoryStartYearEndYearGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidhistorystartyearendyearget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/history/{startYear}/{endYear} | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdMonthlyYearMonthGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidmonthlyyearmonthget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/monthly/{year}/{month} | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdMonthsLastGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidmonthslastget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/months/last | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdTodayGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidtodayget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/today | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdWeekHistoryStartYearEndYearGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidweekhistorystartyearendyearget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/week/history/{startYear}/{endYear} | 
*UnitRevenueApi* | [**ApiV2BrandRevenueUnitCountryIdUnitIdYesterdayGet**](docs/UnitRevenueApi.md#apiv2brandrevenueunitcountryidunitidyesterdayget) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/yesterday | 
*UnitsApi* | [**ApiV2BrandUnitsAllCountryIdGet**](docs/UnitsApi.md#apiv2brandunitsallcountryidget) | **Get** /api/v2/{brand}/units/all/{countryId} | 
*UnitsApi* | [**ApiV2BrandUnitsCountGet**](docs/UnitsApi.md#apiv2brandunitscountget) | **Get** /api/v2/{brand}/units/count | 
*UnitsApi* | [**ApiV2BrandUnitsCountMonthsYearMonthGet**](docs/UnitsApi.md#apiv2brandunitscountmonthsyearmonthget) | **Get** /api/v2/{brand}/units/count/months/{year}/{month} | 
*UnitsApi* | [**ApiV2BrandUnitsCountryIdUnitIdGet**](docs/UnitsApi.md#apiv2brandunitscountryidunitidget) | **Get** /api/v2/{brand}/units/{countryId}/{unitId} | 
*UnitsApi* | [**ApiV2UnitsAllCountryIdGet**](docs/UnitsApi.md#apiv2unitsallcountryidget) | **Get** /api/v2/units/all/{countryId} | 
*UnitsApi* | [**ApiV2UnitsCountGet**](docs/UnitsApi.md#apiv2unitscountget) | **Get** /api/v2/units/count | 
*UnitsApi* | [**ApiV2UnitsCountMonthsYearMonthGet**](docs/UnitsApi.md#apiv2unitscountmonthsyearmonthget) | **Get** /api/v2/units/count/months/{year}/{month} | 

## Documentation For Models

 - [AddressModel](docs/AddressModel.md)
 - [BrandDailyRevenueListModel](docs/BrandDailyRevenueListModel.md)
 - [BrandListDailyRevenueListModel](docs/BrandListDailyRevenueListModel.md)
 - [BrandListTotalUnitCountListModel](docs/BrandListTotalUnitCountListModel.md)
 - [BrandModel](docs/BrandModel.md)
 - [BrandTotalUnitCountListModel](docs/BrandTotalUnitCountListModel.md)
 - [CoordinatesModel](docs/CoordinatesModel.md)
 - [CountryDailyRevenueModel](docs/CountryDailyRevenueModel.md)
 - [CountryDailyRevenueModelBrandData](docs/CountryDailyRevenueModelBrandData.md)
 - [CountryModel](docs/CountryModel.md)
 - [CountryModelBrandData](docs/CountryModelBrandData.md)
 - [DailyMetricsModel](docs/DailyMetricsModel.md)
 - [DailyRevenueListModel](docs/DailyRevenueListModel.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [HouseModel](docs/HouseModel.md)
 - [LocalityModel](docs/LocalityModel.md)
 - [MetricsHistoryModel](docs/MetricsHistoryModel.md)
 - [MetricsModel](docs/MetricsModel.md)
 - [MonthRevenueHistoryModel](docs/MonthRevenueHistoryModel.md)
 - [MonthRevenueHistoryModelBrandData](docs/MonthRevenueHistoryModelBrandData.md)
 - [MonthRevenueModel](docs/MonthRevenueModel.md)
 - [MonthTotalRevenueModel](docs/MonthTotalRevenueModel.md)
 - [MonthTotalRevenueModelBrandData](docs/MonthTotalRevenueModelBrandData.md)
 - [MonthTotalRevenueModelBrandListData](docs/MonthTotalRevenueModelBrandListData.md)
 - [MonthUnitCountHistoryModel](docs/MonthUnitCountHistoryModel.md)
 - [MonthUnitCountHistoryModelBrandData](docs/MonthUnitCountHistoryModelBrandData.md)
 - [MonthUnitCountHistoryModelBrandListData](docs/MonthUnitCountHistoryModelBrandListData.md)
 - [MonthUnitCountModel](docs/MonthUnitCountModel.md)
 - [OrderModel](docs/OrderModel.md)
 - [OrderModelBrandData](docs/OrderModelBrandData.md)
 - [OrderModelBrandListData](docs/OrderModelBrandListData.md)
 - [OrdersCountBySourceStatistics](docs/OrdersCountBySourceStatistics.md)
 - [Period](docs/Period.md)
 - [PeriodicRevenueListModel](docs/PeriodicRevenueListModel.md)
 - [ProductImage](docs/ProductImage.md)
 - [ProductStatistic](docs/ProductStatistic.md)
 - [ProductStatisticsModel](docs/ProductStatisticsModel.md)
 - [ProductStatisticsModelBrandData](docs/ProductStatisticsModelBrandData.md)
 - [Rating](docs/Rating.md)
 - [RatingsModel](docs/RatingsModel.md)
 - [RatingsModelBrandData](docs/RatingsModelBrandData.md)
 - [RatingsModelBrandListData](docs/RatingsModelBrandListData.md)
 - [RevenueHistoryModel](docs/RevenueHistoryModel.md)
 - [RevenueHistoryModelBrandData](docs/RevenueHistoryModelBrandData.md)
 - [RevenueModel](docs/RevenueModel.md)
 - [ScheduleModel](docs/ScheduleModel.md)
 - [StatisticsModel](docs/StatisticsModel.md)
 - [StatisticsModelBrandData](docs/StatisticsModelBrandData.md)
 - [StreetModel](docs/StreetModel.md)
 - [TotalRevenueHistoryModel](docs/TotalRevenueHistoryModel.md)
 - [TotalRevenueHistoryModelBrandData](docs/TotalRevenueHistoryModelBrandData.md)
 - [TotalRevenueHistoryModelBrandListData](docs/TotalRevenueHistoryModelBrandListData.md)
 - [TotalRevenueModel](docs/TotalRevenueModel.md)
 - [TotalUnitCountListModel](docs/TotalUnitCountListModel.md)
 - [UnitCountModel](docs/UnitCountModel.md)
 - [UnitListModel](docs/UnitListModel.md)
 - [UnitListModelBrandData](docs/UnitListModelBrandData.md)
 - [UnitListModelBrandListData](docs/UnitListModelBrandListData.md)
 - [UnitModel](docs/UnitModel.md)
 - [UnitSingleModel](docs/UnitSingleModel.md)
 - [UnitSingleModelBrandData](docs/UnitSingleModelBrandData.md)
 - [WeekRevenueHistoryModel](docs/WeekRevenueHistoryModel.md)
 - [WeekRevenueHistoryModelBrandData](docs/WeekRevenueHistoryModelBrandData.md)
 - [WeekRevenueModel](docs/WeekRevenueModel.md)
 - [WeeklyTopModel](docs/WeeklyTopModel.md)
 - [WeeklyTopModelBrandData](docs/WeeklyTopModelBrandData.md)
 - [WeeklyTopModelBrandListData](docs/WeeklyTopModelBrandListData.md)
 - [WeeklyTopUnitModel](docs/WeeklyTopUnitModel.md)
 - [WorkTimeModel](docs/WorkTimeModel.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

