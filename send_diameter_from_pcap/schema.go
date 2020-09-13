package main

import (
	"net"
	"time"
)

type CCR struct {
	SessionID                     string                               `avp:"Session-Id"`
	UserName                      *string                              `avp:"User-Name"`
	OriginHost                    string                               `avp:"Origin-Host"`
	OriginRealm                   string                               `avp:"Origin-Realm"`
	DestinationHost               string                               `avp:"Destination-Host"`
	DestinationRealm              string                               `avp:"Destination-Realm"`
	AuthApplicationId             uint32                               `avp:"Auth-Application-Id"`
	ServiceContextId              *string                              `avp:"Service-Context-Id"`
	CCRequestType                 int                                  `avp:"CC-Request-Type"`
	CCRequestNumber               int                                  `avp:"CC-Request-Number"`
	OriginStateId                 *uint32                              `avp:"Origin-State-Id"`
	EventTimestamp                *time.Time                           `avp:"Event-Timestamp"`
	SubscriptionId                []*SubscriptionIdData                `avp:"Subscription-Id"`
	MultipleServicesIndicator     *int                                 `avp:"Multiple-Services-Indicator"`
	MultipleServicesCreditControl []*MultipleServicesCreditControlData `avp:"Multiple-Services-Credit-Control"`
	UserEquipmentInfo             *UserEquipmentInfoData               `avp:"User-Equipment-Info"`
	ServiceInformation            *ServiceInformationData              `avp:"Service-Information"`
}

type SubscriptionIdData struct {
	SubscriptionIdType      int    `avp:"Subscription-Id-Type"`
	SubscriptionIdDataValue string `avp:"Subscription-Id-Data"`
}

type MultipleServicesCreditControlData struct {
	RequestedServiceUnit *ServiceUnitData         `avp:"Requested-Service-Unit"`
	UsedServiceUnit      []*ServiceUnitData       `avp:"Used-Service-Unit"`
	GrantedServiceUnit   *ServiceUnitData         `avp:"Granted-Service-Unit"`
	RatingGroup          *uint32                  `avp:"Rating-Group"`
	ServiceIdentifier    *int                     `avp:"Service-Identifier"`
	ValidityTime         *uint32                  `avp:"Validity-Time"`
	ResultCode           *uint32                  `avp:"Result-Code"`
	VolumeQuotaThreshold *uint32                  `avp:"Volume-Quota-Threshold"`
	FinalUnitIndication  *FinalUnitIndicationData `avp:"Final-Unit-Indication"`
}

type FinalUnitIndicationData struct {
	FinalUnitAction *int                `avp:"Final-Unit-Action"`
	RedirectServer  *RedirectServerData `avp:"Redirect-Server"`
}

type RedirectServerData struct {
	RedirectAddressType   *int    `avp:"Redirect-Address-Type"`
	RedirectServerAddress *string `avp:"Redirect-Server-Address"`
}

type ServiceUnitData struct {
	CCInputOctets     *uint64    `avp:"CC-Input-Octets"`
	CCOutputOctets    *uint64    `avp:"CC-Output-Octets"`
	CCTotalOctets     *uint64    `avp:"CC-Total-Octets"`
	CCTime            *uint32    `avp:"CC-Time"`
	TariffChangeUsage *int       `avp:"Tariff-Change-Usage"`
	TariffTimeChange  *time.Time `avp:"Tariff-Time-Change"`
}

type UserEquipmentInfoData struct {
	UserEquipmentInfoType  int    `avp:"User-Equipment-Info-Type"`
	UserEquipmentInfoValue string `avp:"User-Equipment-Info-Value"`
}

type ServiceInformationData struct {
	PSInformation PSInformationData `avp:"PS-Information"`
}

type PSInformationData struct {
	TGPPChargingID               *uint32    `avp:"TGPP-Charging-Id"`
	TGPPPDPType                  *uint32    `avp:"TGPP-PDP-Type"`
	TGPPPPDPAddress              *net.IP    `avp:"PDP-Address"`
	TGPPGPRSNegotiatedQoSProfile *string    `avp:"TGPP-GPRS-Negotiated-QoS-Profile"`
	SGSNAddress                  *net.IP    `avp:"SGSN-Address"`
	GGSNAddress                  *net.IP    `avp:"GGSN-Address"`
	TGPPIMSIMCCMNC               *string    `avp:"TGPP-IMSI-MCC-MNC"`
	TGPPGGSNMCCMNC               *string    `avp:"TGPP-GGSN-MCC-MNC"`
	TGPPSGSNMCCMNC               *string    `avp:"TGPP-SGSN-MCC-MNC"`
	TGPPNSAPI                    *string    `avp:"TGPP-NSAPI"`
	CalledStationId              *string    `avp:"Called-Station-Id"`
	TGPPSelectionMode            *string    `avp:"TGPP-Selection-Mode"`
	TGPPChargingCharacteristics  *string    `avp:"TGPP-Charging-Characteristics"`
	TGPPMSTimeZone               *string    `avp:"TGPP-MS-TimeZone"`
	TGPPUserLocationInfo         *string    `avp:"TGPP-User-Location-Info"`
	TGPPRatType                  *string    `avp:"TGPP-RAT-Type"`
	PDNConnectionChargingID      *uint32    `avp:"PDN-Connection-Charging-Id"`
	ServingNodeType              *int       `avp:"Serving-Node-Type"`
	StartTime                    *time.Time `avp:"Start-Time"`
	ChargingRuleBaseName         *string    `avp:"Charging-Rule-Base-Name"`
}

type CCA struct {
	SessionID                     string                               `avp:"Session-Id"`
	ResultCode                    uint32                               `avp:"Result-Code"`
	OriginHost                    string                               `avp:"Origin-Host"`
	OriginRealm                   string                               `avp:"Origin-Realm"`
	AuthApplicationId             uint32                               `avp:"Auth-Application-Id"`
	CCRequestType                 int                                  `avp:"CC-Request-Type"`
	CCRequestNumber               uint32                               `avp:"CC-Request-Number"`
	CCSessionFailover             int                                  `avp:"CC-Session-Failover"`
	MultipleServicesCreditControl []*MultipleServicesCreditControlData `avp:"Multiple-Services-Credit-Control"`
	CreditControlFailureHandling  int                                  `avp:"Credit-Control-Failure-Handling"`
}
