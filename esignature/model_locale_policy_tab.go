/*
DocuSign REST API

The DocuSign REST API provides you with a powerful, convenient, and simple Web services API for interacting with DocuSign.

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LocalePolicyTab type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalePolicyTab{}

// LocalePolicyTab Allows you to customize locale settings.
type LocalePolicyTab struct {
	// Specifies the address format. Valid values:  - `en_us` - `ja_jp` - `zh_cn_tw` 
	AddressFormat *string `json:"addressFormat,omitempty"`
	// Specifies the type of calendar. Valid values:  - `gregorian` - `japanese` - `buddhist` 
	CalendarType *string `json:"calendarType,omitempty"`
	// The two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code.
	CultureName *string `json:"cultureName,omitempty"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code.  Supported formats:  - `AED` - `AFN` - `ALL` - `AMD` - `ANG` - `AOA` - `ARS` - `AUD` - `AWG` - `AZN` - `BAM` - `BBD` - `BDT` - `BGN` - `BHD` - `BIF` - `BMD` - `BND` - `BOB` - `BOV` - `BRL` - `BSD` - `BTN` - `BWP` - `BYN` - `BYR` - `BZD` - `CAD` - `CDF` - `CHE` - `CHF` - `CHW` - `CLF` - `CLP` - `CNY` - `COP` - `COU` - `CRC` - `CUC` - `CUP` - `CVE` - `CZK` - `DJF` - `DKK` - `DOP` - `DZD` - `EGP` - `ERN` - `ETB` - `EUR` - `FJD` - `FKP` - `GBP` - `GEL` - `GHS` - `GIP` - `GMD` - `GNF` - `GTQ` - `GYD` - `HKD` - `HNL` - `HRK` - `HTG` - `HUF` - `IDR` - `ILS` - `INR` - `IQD` - `IRR` - `ISK` - `JMD` - `JOD` - `JPY` - `KES` - `KGS` - `KHR` - `KMF` - `KPW` - `KRW` - `KWD` - `KYD` - `KZT` - `LAK` - `LBP` - `LKR` - `LRD` - `LSL` - `LYD` - `MAD` - `MDL` - `MGA` - `MKD` - `MMK` - `MNT` - `MOP` - `MRO` - `MUR` - `MVR` - `MWK` - `MXN` - `MXV` - `MYR` - `MZN` - `NAD` - `NGN` - `NIO` - `NOK` - `NPR` - `NZD` - `OMR` - `PAB` - `PEN` - `PGK` - `PHP` - `PKR` - `PLN` - `PYG` - `QAR` - `RON` - `RSD` - `RUB` - `RWF` - `SAR` - `SBD` - `SCR` - `SDG` - `SEK` - `SGD` - `SHP` - `SLL` - `SOS` - `SRD` - `SSP` - `STD` - `SVC` - `SYP` - `SZL` - `THB` - `TJS` - `TMT` - `TND` - `TOP` - `TRY` - `TTD` - `TWD` - `TZS` - `UAH` - `UGX` - `USD` - `USN` - `UYI` - `UYU` - `UZS` - `VEF` - `VND` - `VUV` - `WST` - `XAF` - `XAG` - `XAU` - `XBA` - `XBB` - `XBC` - `XBD` - `XCD` - `XDR` - `XOF` - `XPD` - `XPF` - `XPT` - `XSU` - `XTS` - `XUA` - `XXX` - `YER` - `ZAR` - `ZMW` - `ZWL` 
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Determines how negative currency values are displayed.  In most cases, you should not need to change this value. See [Explicitly define formatting](/docs/esign-rest-api/esign101/concepts/tabs/number-fields/#explicitly-define-formatting).  Valid values:  - `Default`<br>   `0` - `OPar_CSym_1_Comma_234_Comma_567_Period_89_CPar`<br>   `($1,234,567.89)` - `Minus_CSym_1_Comma_234_Comma_567_Period_89`<br>   `-$1,234,567.89` - `Minus_CSym_Space_1_Period_234_Period_567_Comma_89`<br>   `-$ 1.234.567,89` - `CSym_Space_Minus_1_Period_234_Period_567_Comma_89`<br>   `$ -1.234.567,89` - `Minus_1_Period_234_Period_567_Comma_89_Space_CSym`<br>   `-1.234.567,89 $` - `OPar_1_Space_234_Space_567_Comma_89_Space_CSym_CPar`<br>   `(1 234 567,89 $)` - `Minus_1_Space_234_Space_567_Comma_89_Space_CSym`<br>   `-1 234 567,89 $` - `CSym_Minus_1_Quote_234_Quote_567_Period_89`<br>   `$-1'234'567.89` - `Minus_CSym_1_Period_234_Period_567_Comma_89`<br>   `-$1.234.567,89` - `Minus_CSym_1_Comma_234_Comma_567`<br>   `-$1,234,567` - `Minus_CSym_12_Comma_34_Comma_567_Period_89`<br>   `-$12,34,567.89` - `OPar_CSym_Space_1234_Comma_567_Period_89_CPar`<br>   `($ 1234,567.89)` - `CSym_Space_Minus_12_Comma_34_Comma_567_Period_89`<br>   `$ -12,34,567.89` - `CSym_Minus_12_Comma_34_Comma_567_Period_89`<br>   `$-1,234,567.89` - `CSym_Space_Minus_1_Space_234_Space_567_Comma_89`<br>   `$ -1 234 567,89` - `CSym_Space_Minus_1_Space_234_Space_567_Period_89`<br>   `$ -1 234 567.89` - `Minus_CSym_Space_1_Space_234_Space_567_Comma_89`<br>   `-$ 1 234 567,89` - `Minus_1_Space_234_Space_567_Comma_89_CSym`<br>   `-1 234 567,89$` - `Minus_1_Space_234_Space_567_Period_89_Space_CSym`<br>   `-1 234 567.89 $` - `OPar_CSym_1_Period_234_Period_567_CPar`<br>   `(1.234.567)` - `OPar_CSym_1_Comma_234_Comma_567_CPar`<br>   `($1,234,567)` - `Minus_1_Comma_234_Comma_567_Period_89_Space_CSym`<br>   `-1,234,567.89 $` - `Minus_CSym_Space_1_Comma_234_Comma_567_Period_89`<br>   `-$ 1,234,567.89` - `OPar_CSym_Space_1_Period_234_Period_567_Comma_89_CPar`<br>   `($ 1.234.567,89)` - `OPar_CSym_Space_1_Quote_234_Quote_567_Period_89_CPar`<br>   `($ 1'234'567.89)` - `OPar_CSym_Space_1_Space_234_Space_567_Comma_89_CPar`<br>   `($ 1 234 567,89)` - `OPar_CSym_Space_1_Space_234_Space_567_Period_89_CPar`<br>   `($ 1 234 567.89)` - `OPar_CSym_12_Comma_34_Comma_567_Period_89_CPar`<br>   `($12,34,567.89)` - `OPar_CSym_Space_12_Comma_34_Comma_567_Period_89_CPar`<br>   `($ 12,34,567.89)` - `OPar_1_Comma_234_Comma_567_Period_89_Space_CSym_CPar`<br>   `(1,234,567.89 $)` - `OPar_1_Period_234_Period_567_Comma_89_Space_CSym_CPar`<br>   `(1.234.567,89 $)` - `OPar_1_Space_234_Space_567_Comma_89_CSym_CPar`<br>   `(1 234 567,89$)` - `OPar_1_Space_234_Space_567_Period_89_Space_CSym_CPar`<br>   `(1 234 567.89 $)` - `OPar_CSym_Space_1_Comma_234_Comma_567_Period_89_CPar`<br>   `($ 1,234,567.89)` - `Minus_CSym_1_Period_234_Period_567`<br>   `-$ 1.234.567` - `Minus_CSym_Space_1_Quote_234_Quote_567_Period_89`<br>   `-$ 1'234'567.89` - `Minus_CSym_Space_1_Space_234_Space_567_Period_89`<br>   `-$ 1 234 567.89` - `CSym_Minus_1_Comma_234_Comma_567`<br>   `$-1,234,567` - `CSym_Minus_1_Period_234_Period_567`<br>   `$-1.234.567` - `CSym_Space_Minus_1_Quote_234_Quote_567_Period_89`<br>   `$ -1'234'567.89` - `CSym_Space_Minus_1_Comma_234_Comma_567_Period_89`<br>   `$ -1,234,567.89` - `Minus_CSym_Space_12_Comma_34_Comma_567_Period_89`<br>   `-$ 12,34,567.89` - `Minus_1_Period_234_Period_567_Space_CSym`<br>   `-123.456.789 $` - `CSym_Minus_1_Space_234_Space_567_Comma_89`<br>   `$-123 456 789,00` - `Minus_1_Quote_234_Quote_567_Period_89_Space_CSym`<br>   `-123'456'789.00 $` - `CSym_1_Comma_234_Comma_567_Period_89_Minus`<br>   `$123,456,789.00-` - `CSym_Minus_1_Period_234_Period_567_Comma_89`<br>   `$-123.456.789,00` - `OPar_CSym_1_Period_234_Period_567_Comma_89_CPar`<br>   `($123.456.789,00)` - `Minus_CSym_1234_Comma_567_Period_89`<br>   `-$123456,789.00` - `Minus_CSym_1_Space_234_Space_567_Comma_89`<br>   `-$123 456 789,00` 
	CurrencyNegativeFormat *string `json:"currencyNegativeFormat,omitempty"`
	// Determines how positive currency values are displayed.  In most cases, you should not need to change this value. See [Explicitly define formatting](/docs/esign-rest-api/esign101/concepts/tabs/number-fields/#explicitly-define-formatting).  Valid values:   - `Default`<br>   Uses the current locale. - `CSym_1_Comma_234_Comma_567_Period_89`<br>   `$1,234,567.89` - `CSym_Space_1_Period_234_Period_567_Comma_89`<br>   `$ 1.234.567,89` - `Leading_1_Period_234_Period_567_Comma_89_Space_CSym`<br>   `1.234.567,89 $` - `Leading_1_Space_234_Space_567_Comma_89_Space_CSym`<br>   `1 234 567,89 $` - `CSym_Space_1_Quote_234_Quote_567_Period_89`<br>   `$ 1'234'567.89` - `CSym_1_Comma_234_Comma_567`<br>   `$1,234,567` - `CSym_Space_12_Comma_34_Comma_567_Period_89`<br>   `$ 12,34,567.89` - `CSym_12_Comma_34_Comma_567_Period_89`<br>   `$12,34,567.89` - `CSym_Space_1234_Comma_567_Period_89`<br>   `$ 1234,567.89` - `Leading_1_Space_234_Space_567_Period_89_Space_CSym`<br>   `1 234 567.89 $` - `CSym_Space_1_Space_234_Space_567_Comma_89`<br>   `$ 1 234 567,89` - `CSym_Space_1_Space_234_Space_567_Period_89`<br>   `$ 1 234 567.89` - `Leading_1_Space_234_Space_567_Comma_89_CSym`<br>   `1 234 567,89$` - `CSym_1_Period_234_Period_567`<br>   `$1.234.567` - `Leading_1_Comma_234_Comma_567_Period_89_Space_CSym`<br>   `1,234,567. $` (New Armenian) - `CSym_Space_1_Comma_234_Comma_567_Period_89`<br>   `$ 1,234,567.89` (Persian) - `CSym_1_Period_234_Period_567_Comma_89`<br>   `$123.456.789,00` (es-CO) - `Leading_1_Quote_234_Quote_567_Period_89_Space_CSym`<br>   `123'456'789.00 $` (fr-ch) - `CSym_1234_Comma_567_Period_89`<br>   `$123456,789.00` (es-PR) - `Leading_1_Period_234_Period_567_Space_CSym`<br>   `123.456.789 $` - `CSym_1_Space_234_Space_567_Comma_89`<br>   `$123 456 789,00` (en-ZA, es-CR) 
	CurrencyPositiveFormat *string `json:"currencyPositiveFormat,omitempty"`
	// 
	CustomDateFormat *string `json:"customDateFormat,omitempty"`
	// 
	CustomTimeFormat *string `json:"customTimeFormat,omitempty"`
	// Specifies the date format. Valid values:  - `default` <br> used the UI's  - `longformat` <br> use the UI's long format - `dd_mm_yy` <br> dd-MM-yy - `dd_mmm_yy` <br> dd-MMM-yy - `dd_mm_yyyy` <br> dd-MM-yyyy - `dd_mmm_yyyy` <br> dd-MMM-yyyy - `ddmmmmyyyy` <br> dd MMMM yyyy - `ddmmyyyy` <br> dd/MM/yyyy - `ddmmyyyy_de` <br> dd.MM.yyyy - `dmyyyy` <br> d/M/yyyy - `d_m_yyyy` <br> d-M-yyyy - `mmmd_yyyy` <br> MMM d, yyyy - `mmm_dd_yyyy` <br> MMM-dd-yyyy - `mmmmd_yyyy` <br> MMMM d, yyyy - `mm_dd_yyyy` <br> MM-dd-yyyy - `mdyyyy` <br> M/d/yyyy - `yyyy_mmm_dd` <br> yyyy-MMM-dd - `yyyy_mm_dd` <br> yyyy-MM-dd - `yyyymmdd` <br> yyyy/MM/dd - `yyyymd` <br> yyyy/M/d - `custom` <br> Customer set own value - `mmddyyyy` <br> MM/dd/yyyy - `mmddyy` <br> MM/dd/yy - `yyyy_mmmm_d` <br> yyyy MMMM d 
	DateFormat *string `json:"dateFormat,omitempty"`
	// When a user is required to enter their initials, this property specifies how initials are rendered. The examples show the initials for \"William Henry Gates\".   - `first1last1`<br> \"WG\" - `last2`<br> \"GA\" - `first2`<br> \"WI\" - `last2_cjk`<br> first two characters from last name in CJK characters.  <!-- Components/BusinessObjects/Models/ConcealedApiRestModels/localePolicyEnums.cs --> 
	InitialFormat *string `json:"initialFormat,omitempty"`
	// Describes how names are displayed. Valid values:  - `first_middle_last`<br>William Henry Gates - `full`<br>Mr William Henry Gates III - `last_first`<br>Gates William - `lastfirst`<br>GatesWilliam - `last_first_cjk`<br>Gates William only with CJK characters - `lastfirst_cjk`<br>GatesWilliam only with CJK characters  <!-- Web/RestApi/Models/v2_1/localePolicy.cs#L341-L366 --> 
	NameFormat *string `json:"nameFormat,omitempty"`
	// Specifies the time format. Valid values:  - `none`      <br>None - `hh_mm`     <br>hh:mm  - `hhmm`      <br>HH:mm - `hhmmss`    <br>HH:mm:ss - `hhmmsstt`  <br>HH:mm:ss tt - `hhmmtt`    <br> HH:mm tt - `hmm`       <br>h:mm - `hmmss`     <br>h:mm:ss - `hmmsstt`   <br>h:mm:ss tt - `hmmtt`     <br>h:mm tt - `custom`    <br>Customer-set format  <!-- Web/RestApi/Models/v2_1/localePolicy.cs#L501-L546 -->
	TimeFormat *string `json:"timeFormat,omitempty"`
	// Specifies the time zone. Valid values:  - `TZ_01_AfghanistanStandardTime` - `TZ_02_AlaskanStandardTime` - `TZ_03_ArabStandardTime` - `TZ_04_ArabianStandardTime` - `TZ_05_ArabicStandardTime` - `TZ_06_ArgentinaStandardTime` - `TZ_07_AtlanticStandardTime` - `TZ_08_AUS_CentralStandardTime` - `TZ_09_AUS_EasternStandardTime` - `TZ_10_AzerbaijanStandardTime` - `TZ_11_AzoresStandardTime` - `TZ_12_BangladeshStandardTime` - `TZ_13_CanadaCentralStandardTime` - `TZ_14_CapeVerdeStandardTime` - `TZ_15_CaucasusStandardTime` - `TZ_16_CentralAustraliaStandardTime` - `TZ_17_CentralAmericaStandardTime` - `TZ_18_CentralAsiaStandardTime` - `TZ_19_CentralBrazilianStandardTime` - `TZ_20_CentralEuropeStandardTime` - `TZ_21_CentralEuropeanStandardTime` - `TZ_22_CentralPacificStandardTime` - `TZ_23_CentralStandardTime` - `TZ_24_CentralStandardTimeMexico` - `TZ_25_ChinaStandardTime` - `TZ_26_DatelineStandardTime` - `TZ_27_E_AfricaStandardTime` - `TZ_28_E_AustraliaStandardTime` - `TZ_29_E_EuropeStandardTime` - `TZ_30_E_SouthAmericaStandardTime` - `TZ_31_EasternStandardTime` - `TZ_32_EgyptStandardTime` - `TZ_33_EkaterinburgStandardTime` - `TZ_34_FijiStandardTime` - `TZ_35_FLE_StandardTime` - `TZ_36_GeorgianStandardTime` - `TZ_37_GMT_StandardTime` - `TZ_38_GreenlandStandardTime` - `TZ_39_GreenwichStandardTime` - `TZ_40_GTB_StandardTime` - `TZ_41_HawaiianStandardTime` - `TZ_42_IndiaStandardTime` - `TZ_43_IranStandardTime` - `TZ_44_IsraelStandardTime` - `TZ_45_JordanStandardTime` - `TZ_46_KaliningradStandardTime` - `TZ_47_KamchatkaStandardTime` - `TZ_48_KoreaStandardTime` - `TZ_49_MagadanStandardTime` - `TZ_50_MauritiusStandardTime` - `TZ_51_MidAtlanticStandardTime` - `TZ_52_MiddleEastStandardTime` - `TZ_53_MontevideoStandardTime` - `TZ_54_MoroccoStandardTime` - `TZ_55_MountainStandardTime` - `TZ_56_MountainStandardTimeMMexico` - `TZ_57_MyanmarStandardTime` - `TZ_58_N_CentralAsiaStandardTime` - `TZ_59_NamibiaStandardTime` - `TZ_60_NepalStandardTime` - `TZ_61_NewZealandStandardTime` - `TZ_62_NewfoundlandStandardTime` - `TZ_63_NorthAsiaEastStandardTime` - `TZ_64_NorthAsiaStandardTime` - `TZ_65_PacificSAStandardTime` - `TZ_66_PacificStandardTime` - `TZ_67_PacificStandardTimeMexico` - `TZ_68_PakistanStandardTime` - `TZ_69_ParaguayStandardTime` - `TZ_70_RomanceStandardTime` - `TZ_71_RussianStandardTime` - `TZ_72_SAEasternStandardTime` - `TZ_73_SAPacificStandardTime` - `TZ_74_SAWesternStandardTime` - `TZ_75_SamoaStandardTime` - `TZ_76_SE_AsiaStandardTime` - `TZ_77_SingaporeStandardTime` - `TZ_78_SouthAfricaStandardTime` - `TZ_79_SriLankaStandardTime` - `TZ_80_SyriaStandardTime` - `TZ_81_TaipeiStandardTime` - `TZ_82_TasmaniaStandardTime` - `TZ_83_TokyoStandardTime` - `TZ_84_TongaStandardTime` - `TZ_85_TurkeyStandardTime` - `TZ_86_UlaanbaatarStandardTime` - `TZ_87_US_EasternStandardTime` - `TZ_88_USMountainStandardTime` - `TZ_89_VenezuelaStandardTime` - `TZ_90_VladivostokStandardTime` - `TZ_91_W_AustraliaStandardTime` - `TZ_92_W_CentralAfricaStandardTime` - `TZ_93_W_EuropeStandardTime` - `TZ_94_WestAsiaStandardTime` - `TZ_95_WestPacificStandardTime` - `TZ_96_YakutskStandardTime` 
	TimeZone *string `json:"timeZone,omitempty"`
	// When **true,** use the long currency format for the locale.
	UseLongCurrencyFormat *string `json:"useLongCurrencyFormat,omitempty"`
}

// NewLocalePolicyTab instantiates a new LocalePolicyTab object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalePolicyTab() *LocalePolicyTab {
	this := LocalePolicyTab{}
	return &this
}

// NewLocalePolicyTabWithDefaults instantiates a new LocalePolicyTab object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalePolicyTabWithDefaults() *LocalePolicyTab {
	this := LocalePolicyTab{}
	return &this
}

// GetAddressFormat returns the AddressFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetAddressFormat() string {
	if o == nil || IsNil(o.AddressFormat) {
		var ret string
		return ret
	}
	return *o.AddressFormat
}

// GetAddressFormatOk returns a tuple with the AddressFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetAddressFormatOk() (*string, bool) {
	if o == nil || IsNil(o.AddressFormat) {
		return nil, false
	}
	return o.AddressFormat, true
}

// HasAddressFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasAddressFormat() bool {
	if o != nil && !IsNil(o.AddressFormat) {
		return true
	}

	return false
}

// SetAddressFormat gets a reference to the given string and assigns it to the AddressFormat field.
func (o *LocalePolicyTab) SetAddressFormat(v string) {
	o.AddressFormat = &v
}

// GetCalendarType returns the CalendarType field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCalendarType() string {
	if o == nil || IsNil(o.CalendarType) {
		var ret string
		return ret
	}
	return *o.CalendarType
}

// GetCalendarTypeOk returns a tuple with the CalendarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCalendarTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CalendarType) {
		return nil, false
	}
	return o.CalendarType, true
}

// HasCalendarType returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCalendarType() bool {
	if o != nil && !IsNil(o.CalendarType) {
		return true
	}

	return false
}

// SetCalendarType gets a reference to the given string and assigns it to the CalendarType field.
func (o *LocalePolicyTab) SetCalendarType(v string) {
	o.CalendarType = &v
}

// GetCultureName returns the CultureName field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCultureName() string {
	if o == nil || IsNil(o.CultureName) {
		var ret string
		return ret
	}
	return *o.CultureName
}

// GetCultureNameOk returns a tuple with the CultureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCultureNameOk() (*string, bool) {
	if o == nil || IsNil(o.CultureName) {
		return nil, false
	}
	return o.CultureName, true
}

// HasCultureName returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCultureName() bool {
	if o != nil && !IsNil(o.CultureName) {
		return true
	}

	return false
}

// SetCultureName gets a reference to the given string and assigns it to the CultureName field.
func (o *LocalePolicyTab) SetCultureName(v string) {
	o.CultureName = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *LocalePolicyTab) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencyNegativeFormat returns the CurrencyNegativeFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCurrencyNegativeFormat() string {
	if o == nil || IsNil(o.CurrencyNegativeFormat) {
		var ret string
		return ret
	}
	return *o.CurrencyNegativeFormat
}

// GetCurrencyNegativeFormatOk returns a tuple with the CurrencyNegativeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCurrencyNegativeFormatOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyNegativeFormat) {
		return nil, false
	}
	return o.CurrencyNegativeFormat, true
}

// HasCurrencyNegativeFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCurrencyNegativeFormat() bool {
	if o != nil && !IsNil(o.CurrencyNegativeFormat) {
		return true
	}

	return false
}

// SetCurrencyNegativeFormat gets a reference to the given string and assigns it to the CurrencyNegativeFormat field.
func (o *LocalePolicyTab) SetCurrencyNegativeFormat(v string) {
	o.CurrencyNegativeFormat = &v
}

// GetCurrencyPositiveFormat returns the CurrencyPositiveFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCurrencyPositiveFormat() string {
	if o == nil || IsNil(o.CurrencyPositiveFormat) {
		var ret string
		return ret
	}
	return *o.CurrencyPositiveFormat
}

// GetCurrencyPositiveFormatOk returns a tuple with the CurrencyPositiveFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCurrencyPositiveFormatOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyPositiveFormat) {
		return nil, false
	}
	return o.CurrencyPositiveFormat, true
}

// HasCurrencyPositiveFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCurrencyPositiveFormat() bool {
	if o != nil && !IsNil(o.CurrencyPositiveFormat) {
		return true
	}

	return false
}

// SetCurrencyPositiveFormat gets a reference to the given string and assigns it to the CurrencyPositiveFormat field.
func (o *LocalePolicyTab) SetCurrencyPositiveFormat(v string) {
	o.CurrencyPositiveFormat = &v
}

// GetCustomDateFormat returns the CustomDateFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCustomDateFormat() string {
	if o == nil || IsNil(o.CustomDateFormat) {
		var ret string
		return ret
	}
	return *o.CustomDateFormat
}

// GetCustomDateFormatOk returns a tuple with the CustomDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCustomDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.CustomDateFormat) {
		return nil, false
	}
	return o.CustomDateFormat, true
}

// HasCustomDateFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCustomDateFormat() bool {
	if o != nil && !IsNil(o.CustomDateFormat) {
		return true
	}

	return false
}

// SetCustomDateFormat gets a reference to the given string and assigns it to the CustomDateFormat field.
func (o *LocalePolicyTab) SetCustomDateFormat(v string) {
	o.CustomDateFormat = &v
}

// GetCustomTimeFormat returns the CustomTimeFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetCustomTimeFormat() string {
	if o == nil || IsNil(o.CustomTimeFormat) {
		var ret string
		return ret
	}
	return *o.CustomTimeFormat
}

// GetCustomTimeFormatOk returns a tuple with the CustomTimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetCustomTimeFormatOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTimeFormat) {
		return nil, false
	}
	return o.CustomTimeFormat, true
}

// HasCustomTimeFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasCustomTimeFormat() bool {
	if o != nil && !IsNil(o.CustomTimeFormat) {
		return true
	}

	return false
}

// SetCustomTimeFormat gets a reference to the given string and assigns it to the CustomTimeFormat field.
func (o *LocalePolicyTab) SetCustomTimeFormat(v string) {
	o.CustomTimeFormat = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetDateFormat() string {
	if o == nil || IsNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormat) {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasDateFormat() bool {
	if o != nil && !IsNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *LocalePolicyTab) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetInitialFormat returns the InitialFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetInitialFormat() string {
	if o == nil || IsNil(o.InitialFormat) {
		var ret string
		return ret
	}
	return *o.InitialFormat
}

// GetInitialFormatOk returns a tuple with the InitialFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetInitialFormatOk() (*string, bool) {
	if o == nil || IsNil(o.InitialFormat) {
		return nil, false
	}
	return o.InitialFormat, true
}

// HasInitialFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasInitialFormat() bool {
	if o != nil && !IsNil(o.InitialFormat) {
		return true
	}

	return false
}

// SetInitialFormat gets a reference to the given string and assigns it to the InitialFormat field.
func (o *LocalePolicyTab) SetInitialFormat(v string) {
	o.InitialFormat = &v
}

// GetNameFormat returns the NameFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetNameFormat() string {
	if o == nil || IsNil(o.NameFormat) {
		var ret string
		return ret
	}
	return *o.NameFormat
}

// GetNameFormatOk returns a tuple with the NameFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetNameFormatOk() (*string, bool) {
	if o == nil || IsNil(o.NameFormat) {
		return nil, false
	}
	return o.NameFormat, true
}

// HasNameFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasNameFormat() bool {
	if o != nil && !IsNil(o.NameFormat) {
		return true
	}

	return false
}

// SetNameFormat gets a reference to the given string and assigns it to the NameFormat field.
func (o *LocalePolicyTab) SetNameFormat(v string) {
	o.NameFormat = &v
}

// GetTimeFormat returns the TimeFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetTimeFormat() string {
	if o == nil || IsNil(o.TimeFormat) {
		var ret string
		return ret
	}
	return *o.TimeFormat
}

// GetTimeFormatOk returns a tuple with the TimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetTimeFormatOk() (*string, bool) {
	if o == nil || IsNil(o.TimeFormat) {
		return nil, false
	}
	return o.TimeFormat, true
}

// HasTimeFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasTimeFormat() bool {
	if o != nil && !IsNil(o.TimeFormat) {
		return true
	}

	return false
}

// SetTimeFormat gets a reference to the given string and assigns it to the TimeFormat field.
func (o *LocalePolicyTab) SetTimeFormat(v string) {
	o.TimeFormat = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *LocalePolicyTab) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetUseLongCurrencyFormat returns the UseLongCurrencyFormat field value if set, zero value otherwise.
func (o *LocalePolicyTab) GetUseLongCurrencyFormat() string {
	if o == nil || IsNil(o.UseLongCurrencyFormat) {
		var ret string
		return ret
	}
	return *o.UseLongCurrencyFormat
}

// GetUseLongCurrencyFormatOk returns a tuple with the UseLongCurrencyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalePolicyTab) GetUseLongCurrencyFormatOk() (*string, bool) {
	if o == nil || IsNil(o.UseLongCurrencyFormat) {
		return nil, false
	}
	return o.UseLongCurrencyFormat, true
}

// HasUseLongCurrencyFormat returns a boolean if a field has been set.
func (o *LocalePolicyTab) HasUseLongCurrencyFormat() bool {
	if o != nil && !IsNil(o.UseLongCurrencyFormat) {
		return true
	}

	return false
}

// SetUseLongCurrencyFormat gets a reference to the given string and assigns it to the UseLongCurrencyFormat field.
func (o *LocalePolicyTab) SetUseLongCurrencyFormat(v string) {
	o.UseLongCurrencyFormat = &v
}

func (o LocalePolicyTab) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalePolicyTab) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFormat) {
		toSerialize["addressFormat"] = o.AddressFormat
	}
	if !IsNil(o.CalendarType) {
		toSerialize["calendarType"] = o.CalendarType
	}
	if !IsNil(o.CultureName) {
		toSerialize["cultureName"] = o.CultureName
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencyNegativeFormat) {
		toSerialize["currencyNegativeFormat"] = o.CurrencyNegativeFormat
	}
	if !IsNil(o.CurrencyPositiveFormat) {
		toSerialize["currencyPositiveFormat"] = o.CurrencyPositiveFormat
	}
	if !IsNil(o.CustomDateFormat) {
		toSerialize["customDateFormat"] = o.CustomDateFormat
	}
	if !IsNil(o.CustomTimeFormat) {
		toSerialize["customTimeFormat"] = o.CustomTimeFormat
	}
	if !IsNil(o.DateFormat) {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if !IsNil(o.InitialFormat) {
		toSerialize["initialFormat"] = o.InitialFormat
	}
	if !IsNil(o.NameFormat) {
		toSerialize["nameFormat"] = o.NameFormat
	}
	if !IsNil(o.TimeFormat) {
		toSerialize["timeFormat"] = o.TimeFormat
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.UseLongCurrencyFormat) {
		toSerialize["useLongCurrencyFormat"] = o.UseLongCurrencyFormat
	}
	return toSerialize, nil
}

type NullableLocalePolicyTab struct {
	value *LocalePolicyTab
	isSet bool
}

func (v NullableLocalePolicyTab) Get() *LocalePolicyTab {
	return v.value
}

func (v *NullableLocalePolicyTab) Set(val *LocalePolicyTab) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalePolicyTab) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalePolicyTab) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalePolicyTab(val *LocalePolicyTab) *NullableLocalePolicyTab {
	return &NullableLocalePolicyTab{value: val, isSet: true}
}

func (v NullableLocalePolicyTab) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalePolicyTab) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


