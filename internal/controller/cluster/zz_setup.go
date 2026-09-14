// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aigatewaymcpservice "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/aigatewaymcpservice"
	aigatewaymodelproviderservice "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/aigatewaymodelproviderservice"
	aigatewaymodelservice "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/aigatewaymodelservice"
	aisearchendpoint "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/aisearchendpoint"
	aisearchindex "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/aisearchindex"
	knowledgeassistant "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/knowledgeassistant"
	knowledgeassistantknowledgesource "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/knowledgeassistantknowledgesource"
	supervisoragent "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/supervisoragent"
	supervisoragenttool "github.com/buzz-IT/provider-databricks/internal/controller/cluster/ai/supervisoragenttool"
	app "github.com/buzz-IT/provider-databricks/internal/controller/cluster/apps/app"
	appspace "github.com/buzz-IT/provider-databricks/internal/controller/cluster/apps/appspace"
	appssettingscustomtemplate "github.com/buzz-IT/provider-databricks/internal/controller/cluster/apps/appssettingscustomtemplate"
	customappintegration "github.com/buzz-IT/provider-databricks/internal/controller/cluster/apps/customappintegration"
	budgetpolicy "github.com/buzz-IT/provider-databricks/internal/controller/cluster/billing/budgetpolicy"
	cluster "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/cluster"
	clusterpolicy "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/clusterpolicy"
	instancepool "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/instancepool"
	job "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/job"
	library "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/library"
	pipeline "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/pipeline"
	sandbox "github.com/buzz-IT/provider-databricks/internal/controller/cluster/compute/sandbox"
	databasedatabasecatalog "github.com/buzz-IT/provider-databricks/internal/controller/cluster/databases/databasedatabasecatalog"
	databaseinstance "github.com/buzz-IT/provider-databricks/internal/controller/cluster/databases/databaseinstance"
	databasesynceddatabasetable "github.com/buzz-IT/provider-databricks/internal/controller/cluster/databases/databasesynceddatabasetable"
	instanceprofile "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/instanceprofile"
	mwscredentials "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwscredentials"
	mwscustomermanagedkeys "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwscustomermanagedkeys"
	mwsnccbinding "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsnccbinding"
	mwsnccprivateendpointrule "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsnccprivateendpointrule"
	mwsnetworkconnectivityconfig "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsnetworkconnectivityconfig"
	mwsnetworks "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsnetworks"
	mwsprivateaccesssettings "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsprivateaccesssettings"
	mwsstorageconfigurations "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsstorageconfigurations"
	mwsvpcendpoint "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsvpcendpoint"
	mwsworkspaces "github.com/buzz-IT/provider-databricks/internal/controller/cluster/deployment/mwsworkspaces"
	disasterrecoveryfailovergroup "github.com/buzz-IT/provider-databricks/internal/controller/cluster/dr/disasterrecoveryfailovergroup"
	disasterrecoverystableurl "github.com/buzz-IT/provider-databricks/internal/controller/cluster/dr/disasterrecoverystableurl"
	environmentsdefaultworkspacebaseenvironment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/envs/environmentsdefaultworkspacebaseenvironment"
	environmentsworkspacebaseenvironment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/envs/environmentsworkspacebaseenvironment"
	budget "github.com/buzz-IT/provider-databricks/internal/controller/cluster/finops/budget"
	dataclassificationcatalogconfig "github.com/buzz-IT/provider-databricks/internal/controller/cluster/governance/dataclassificationcatalogconfig"
	mwslogdelivery "github.com/buzz-IT/provider-databricks/internal/controller/cluster/log/mwslogdelivery"
	featureengineeringfeature "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/featureengineeringfeature"
	featureengineeringkafkaconfig "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/featureengineeringkafkaconfig"
	featureengineeringmaterializedfeature "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/featureengineeringmaterializedfeature"
	materializedfeaturesfeaturetag "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/materializedfeaturesfeaturetag"
	mlflowexperiment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/mlflowexperiment"
	mlflowmodel "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/mlflowmodel"
	mlflowwebhook "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mlflow/mlflowwebhook"
	vectorsearchendpoint "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mosaic/vectorsearchendpoint"
	vectorsearchindex "github.com/buzz-IT/provider-databricks/internal/controller/cluster/mosaic/vectorsearchindex"
	accountfederationpolicy "github.com/buzz-IT/provider-databricks/internal/controller/cluster/oauth/accountfederationpolicy"
	serviceprincipalfederationpolicy "github.com/buzz-IT/provider-databricks/internal/controller/cluster/oauth/serviceprincipalfederationpolicy"
	postgresbranch "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgresbranch"
	postgrescatalog "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgrescatalog"
	postgrescdfconfig "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgrescdfconfig"
	postgresdataapi "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgresdataapi"
	postgresdatabase "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgresdatabase"
	postgresendpoint "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgresendpoint"
	postgresproject "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgresproject"
	postgresrole "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgresrole"
	postgressnapshotschedule "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgressnapshotschedule"
	postgressyncedtable "github.com/buzz-IT/provider-databricks/internal/controller/cluster/postgres/postgressyncedtable"
	providerconfig "github.com/buzz-IT/provider-databricks/internal/controller/cluster/providerconfig"
	accesscontrolruleset "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/accesscontrolruleset"
	accountiamdirectgroupmemberv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/accountiamdirectgroupmemberv2"
	accountiamgroupv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/accountiamgroupv2"
	accountiamserviceprincipalv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/accountiamserviceprincipalv2"
	accountiamuserv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/accountiamuserv2"
	accountiamworkspaceassignmentv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/accountiamworkspaceassignmentv2"
	entitlements "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/entitlements"
	group "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/group"
	groupinstanceprofile "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/groupinstanceprofile"
	groupmember "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/groupmember"
	grouprole "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/grouprole"
	ipaccesslist "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/ipaccesslist"
	mwspermissionassignment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/mwspermissionassignment"
	obotoken "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/obotoken"
	permissionassignment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/permissionassignment"
	permissions "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/permissions"
	secret "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/secret"
	secretacl "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/secretacl"
	secretscope "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/secretscope"
	serviceprincipal "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/serviceprincipal"
	serviceprincipalrole "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/serviceprincipalrole"
	serviceprincipalsecret "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/serviceprincipalsecret"
	sqlpermissions "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/sqlpermissions"
	token "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/token"
	user "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/user"
	userinstanceprofile "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/userinstanceprofile"
	userrole "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/userrole"
	workspaceiamdirectgroupmemberv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/workspaceiamdirectgroupmemberv2"
	workspaceiamgroupv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/workspaceiamgroupv2"
	workspaceiamserviceprincipalv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/workspaceiamserviceprincipalv2"
	workspaceiamuserv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/workspaceiamuserv2"
	workspaceiamworkspaceassignmentv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/workspaceiamworkspaceassignmentv2"
	workspaceiamworkspaceidentitydetailv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/security/workspaceiamworkspaceidentitydetailv2"
	endpoint "github.com/buzz-IT/provider-databricks/internal/controller/cluster/serving/endpoint"
	modelserving "github.com/buzz-IT/provider-databricks/internal/controller/cluster/serving/modelserving"
	modelservingprovisionedthroughput "github.com/buzz-IT/provider-databricks/internal/controller/cluster/serving/modelservingprovisionedthroughput"
	onlinestore "github.com/buzz-IT/provider-databricks/internal/controller/cluster/serving/onlinestore"
	accountnetworkpolicy "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/accountnetworkpolicy"
	accountsettinguserpreferencev2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/accountsettinguserpreferencev2"
	accountsettingv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/accountsettingv2"
	aibidashboardembeddingaccesspolicysetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/aibidashboardembeddingaccesspolicysetting"
	aibidashboardembeddingapproveddomainssetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/aibidashboardembeddingapproveddomainssetting"
	automaticclusterupdateworkspacesetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/automaticclusterupdateworkspacesetting"
	compliancesecurityprofileworkspacesetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/compliancesecurityprofileworkspacesetting"
	defaultnamespacesetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/defaultnamespacesetting"
	disablelegacyaccesssetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/disablelegacyaccesssetting"
	disablelegacydbfssetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/disablelegacydbfssetting"
	disablelegacyfeaturessetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/disablelegacyfeaturessetting"
	enhancedsecuritymonitoringworkspacesetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/enhancedsecuritymonitoringworkspacesetting"
	restrictworkspaceadminssetting "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/restrictworkspaceadminssetting"
	workspacenetworkoption "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/workspacenetworkoption"
	workspacesettingv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/settings/workspacesettingv2"
	provider "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sharing/provider"
	recipient "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sharing/recipient"
	share "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sharing/share"
	alert "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/alert"
	alertv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/alertv2"
	dashboard "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/dashboard"
	query "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/query"
	sqlalert "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqlalert"
	sqldashboard "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqldashboard"
	sqlendpoint "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqlendpoint"
	sqlglobalconfig "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqlglobalconfig"
	sqlquery "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqlquery"
	sqlvisualization "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqlvisualization"
	sqlwidget "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/sqlwidget"
	warehousesdefaultwarehouseoverride "github.com/buzz-IT/provider-databricks/internal/controller/cluster/sql/warehousesdefaultwarehouseoverride"
	awss3mount "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/awss3mount"
	azureadlsgen1mount "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/azureadlsgen1mount"
	azureadlsgen2mount "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/azureadlsgen2mount"
	azureblobmount "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/azureblobmount"
	dbfsfile "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/dbfsfile"
	file "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/file"
	mount "github.com/buzz-IT/provider-databricks/internal/controller/cluster/storage/mount"
	tagpolicy "github.com/buzz-IT/provider-databricks/internal/controller/cluster/tags/tagpolicy"
	workspaceentitytagassignment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/tags/workspaceentitytagassignment"
	artifactallowlist "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/artifactallowlist"
	catalog "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/catalog"
	catalogworkspacebinding "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/catalogworkspacebinding"
	connection "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/connection"
	credential "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/credential"
	dataqualitymonitor "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/dataqualitymonitor"
	dataqualityrefresh "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/dataqualityrefresh"
	domain "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/domain"
	entitytagassignment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/entitytagassignment"
	externallocation "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/externallocation"
	externalmetadata "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/externalmetadata"
	grant "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/grant"
	grantmap "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/grantmap"
	lakehousemonitor "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/lakehousemonitor"
	metastore "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/metastore"
	metastoreassignment "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/metastoreassignment"
	metastoredataaccess "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/metastoredataaccess"
	onlinetable "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/onlinetable"
	policyinfo "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/policyinfo"
	qualitymonitor "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/qualitymonitor"
	qualitymonitorv2 "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/qualitymonitorv2"
	registeredmodel "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/registeredmodel"
	rfaaccessrequestdestinations "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/rfaaccessrequestdestinations"
	schema "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/schema"
	secretuc "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/secretuc"
	sqltable "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/sqltable"
	storagecredential "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/storagecredential"
	systemschema "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/systemschema"
	table "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/table"
	volume "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/volume"
	workspacebinding "github.com/buzz-IT/provider-databricks/internal/controller/cluster/unity/workspacebinding"
	directory "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/directory"
	gitcredential "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/gitcredential"
	globalinitscript "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/globalinitscript"
	notebook "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/notebook"
	notificationdestination "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/notificationdestination"
	repo "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/repo"
	workspaceconf "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/workspaceconf"
	workspacefile "github.com/buzz-IT/provider-databricks/internal/controller/cluster/workspace/workspacefile"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aigatewaymcpservice.Setup,
		aigatewaymodelproviderservice.Setup,
		aigatewaymodelservice.Setup,
		aisearchendpoint.Setup,
		aisearchindex.Setup,
		knowledgeassistant.Setup,
		knowledgeassistantknowledgesource.Setup,
		supervisoragent.Setup,
		supervisoragenttool.Setup,
		app.Setup,
		appspace.Setup,
		appssettingscustomtemplate.Setup,
		customappintegration.Setup,
		budgetpolicy.Setup,
		cluster.Setup,
		clusterpolicy.Setup,
		instancepool.Setup,
		job.Setup,
		library.Setup,
		pipeline.Setup,
		sandbox.Setup,
		databasedatabasecatalog.Setup,
		databaseinstance.Setup,
		databasesynceddatabasetable.Setup,
		instanceprofile.Setup,
		mwscredentials.Setup,
		mwscustomermanagedkeys.Setup,
		mwsnccbinding.Setup,
		mwsnccprivateendpointrule.Setup,
		mwsnetworkconnectivityconfig.Setup,
		mwsnetworks.Setup,
		mwsprivateaccesssettings.Setup,
		mwsstorageconfigurations.Setup,
		mwsvpcendpoint.Setup,
		mwsworkspaces.Setup,
		disasterrecoveryfailovergroup.Setup,
		disasterrecoverystableurl.Setup,
		environmentsdefaultworkspacebaseenvironment.Setup,
		environmentsworkspacebaseenvironment.Setup,
		budget.Setup,
		dataclassificationcatalogconfig.Setup,
		mwslogdelivery.Setup,
		featureengineeringfeature.Setup,
		featureengineeringkafkaconfig.Setup,
		featureengineeringmaterializedfeature.Setup,
		materializedfeaturesfeaturetag.Setup,
		mlflowexperiment.Setup,
		mlflowmodel.Setup,
		mlflowwebhook.Setup,
		vectorsearchendpoint.Setup,
		vectorsearchindex.Setup,
		accountfederationpolicy.Setup,
		serviceprincipalfederationpolicy.Setup,
		postgresbranch.Setup,
		postgrescatalog.Setup,
		postgrescdfconfig.Setup,
		postgresdataapi.Setup,
		postgresdatabase.Setup,
		postgresendpoint.Setup,
		postgresproject.Setup,
		postgresrole.Setup,
		postgressnapshotschedule.Setup,
		postgressyncedtable.Setup,
		providerconfig.Setup,
		accesscontrolruleset.Setup,
		accountiamdirectgroupmemberv2.Setup,
		accountiamgroupv2.Setup,
		accountiamserviceprincipalv2.Setup,
		accountiamuserv2.Setup,
		accountiamworkspaceassignmentv2.Setup,
		entitlements.Setup,
		group.Setup,
		groupinstanceprofile.Setup,
		groupmember.Setup,
		grouprole.Setup,
		ipaccesslist.Setup,
		mwspermissionassignment.Setup,
		obotoken.Setup,
		permissionassignment.Setup,
		permissions.Setup,
		secret.Setup,
		secretacl.Setup,
		secretscope.Setup,
		serviceprincipal.Setup,
		serviceprincipalrole.Setup,
		serviceprincipalsecret.Setup,
		sqlpermissions.Setup,
		token.Setup,
		user.Setup,
		userinstanceprofile.Setup,
		userrole.Setup,
		workspaceiamdirectgroupmemberv2.Setup,
		workspaceiamgroupv2.Setup,
		workspaceiamserviceprincipalv2.Setup,
		workspaceiamuserv2.Setup,
		workspaceiamworkspaceassignmentv2.Setup,
		workspaceiamworkspaceidentitydetailv2.Setup,
		endpoint.Setup,
		modelserving.Setup,
		modelservingprovisionedthroughput.Setup,
		onlinestore.Setup,
		accountnetworkpolicy.Setup,
		accountsettinguserpreferencev2.Setup,
		accountsettingv2.Setup,
		aibidashboardembeddingaccesspolicysetting.Setup,
		aibidashboardembeddingapproveddomainssetting.Setup,
		automaticclusterupdateworkspacesetting.Setup,
		compliancesecurityprofileworkspacesetting.Setup,
		defaultnamespacesetting.Setup,
		disablelegacyaccesssetting.Setup,
		disablelegacydbfssetting.Setup,
		disablelegacyfeaturessetting.Setup,
		enhancedsecuritymonitoringworkspacesetting.Setup,
		restrictworkspaceadminssetting.Setup,
		workspacenetworkoption.Setup,
		workspacesettingv2.Setup,
		provider.Setup,
		recipient.Setup,
		share.Setup,
		alert.Setup,
		alertv2.Setup,
		dashboard.Setup,
		query.Setup,
		sqlalert.Setup,
		sqldashboard.Setup,
		sqlendpoint.Setup,
		sqlglobalconfig.Setup,
		sqlquery.Setup,
		sqlvisualization.Setup,
		sqlwidget.Setup,
		warehousesdefaultwarehouseoverride.Setup,
		awss3mount.Setup,
		azureadlsgen1mount.Setup,
		azureadlsgen2mount.Setup,
		azureblobmount.Setup,
		dbfsfile.Setup,
		file.Setup,
		mount.Setup,
		tagpolicy.Setup,
		workspaceentitytagassignment.Setup,
		artifactallowlist.Setup,
		catalog.Setup,
		catalogworkspacebinding.Setup,
		connection.Setup,
		credential.Setup,
		dataqualitymonitor.Setup,
		dataqualityrefresh.Setup,
		domain.Setup,
		entitytagassignment.Setup,
		externallocation.Setup,
		externalmetadata.Setup,
		grant.Setup,
		grantmap.Setup,
		lakehousemonitor.Setup,
		metastore.Setup,
		metastoreassignment.Setup,
		metastoredataaccess.Setup,
		onlinetable.Setup,
		policyinfo.Setup,
		qualitymonitor.Setup,
		qualitymonitorv2.Setup,
		registeredmodel.Setup,
		rfaaccessrequestdestinations.Setup,
		schema.Setup,
		secretuc.Setup,
		sqltable.Setup,
		storagecredential.Setup,
		systemschema.Setup,
		table.Setup,
		volume.Setup,
		workspacebinding.Setup,
		directory.Setup,
		gitcredential.Setup,
		globalinitscript.Setup,
		notebook.Setup,
		notificationdestination.Setup,
		repo.Setup,
		workspaceconf.Setup,
		workspacefile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aigatewaymcpservice.SetupGated,
		aigatewaymodelproviderservice.SetupGated,
		aigatewaymodelservice.SetupGated,
		aisearchendpoint.SetupGated,
		aisearchindex.SetupGated,
		knowledgeassistant.SetupGated,
		knowledgeassistantknowledgesource.SetupGated,
		supervisoragent.SetupGated,
		supervisoragenttool.SetupGated,
		app.SetupGated,
		appspace.SetupGated,
		appssettingscustomtemplate.SetupGated,
		customappintegration.SetupGated,
		budgetpolicy.SetupGated,
		cluster.SetupGated,
		clusterpolicy.SetupGated,
		instancepool.SetupGated,
		job.SetupGated,
		library.SetupGated,
		pipeline.SetupGated,
		sandbox.SetupGated,
		databasedatabasecatalog.SetupGated,
		databaseinstance.SetupGated,
		databasesynceddatabasetable.SetupGated,
		instanceprofile.SetupGated,
		mwscredentials.SetupGated,
		mwscustomermanagedkeys.SetupGated,
		mwsnccbinding.SetupGated,
		mwsnccprivateendpointrule.SetupGated,
		mwsnetworkconnectivityconfig.SetupGated,
		mwsnetworks.SetupGated,
		mwsprivateaccesssettings.SetupGated,
		mwsstorageconfigurations.SetupGated,
		mwsvpcendpoint.SetupGated,
		mwsworkspaces.SetupGated,
		disasterrecoveryfailovergroup.SetupGated,
		disasterrecoverystableurl.SetupGated,
		environmentsdefaultworkspacebaseenvironment.SetupGated,
		environmentsworkspacebaseenvironment.SetupGated,
		budget.SetupGated,
		dataclassificationcatalogconfig.SetupGated,
		mwslogdelivery.SetupGated,
		featureengineeringfeature.SetupGated,
		featureengineeringkafkaconfig.SetupGated,
		featureengineeringmaterializedfeature.SetupGated,
		materializedfeaturesfeaturetag.SetupGated,
		mlflowexperiment.SetupGated,
		mlflowmodel.SetupGated,
		mlflowwebhook.SetupGated,
		vectorsearchendpoint.SetupGated,
		vectorsearchindex.SetupGated,
		accountfederationpolicy.SetupGated,
		serviceprincipalfederationpolicy.SetupGated,
		postgresbranch.SetupGated,
		postgrescatalog.SetupGated,
		postgrescdfconfig.SetupGated,
		postgresdataapi.SetupGated,
		postgresdatabase.SetupGated,
		postgresendpoint.SetupGated,
		postgresproject.SetupGated,
		postgresrole.SetupGated,
		postgressnapshotschedule.SetupGated,
		postgressyncedtable.SetupGated,
		providerconfig.SetupGated,
		accesscontrolruleset.SetupGated,
		accountiamdirectgroupmemberv2.SetupGated,
		accountiamgroupv2.SetupGated,
		accountiamserviceprincipalv2.SetupGated,
		accountiamuserv2.SetupGated,
		accountiamworkspaceassignmentv2.SetupGated,
		entitlements.SetupGated,
		group.SetupGated,
		groupinstanceprofile.SetupGated,
		groupmember.SetupGated,
		grouprole.SetupGated,
		ipaccesslist.SetupGated,
		mwspermissionassignment.SetupGated,
		obotoken.SetupGated,
		permissionassignment.SetupGated,
		permissions.SetupGated,
		secret.SetupGated,
		secretacl.SetupGated,
		secretscope.SetupGated,
		serviceprincipal.SetupGated,
		serviceprincipalrole.SetupGated,
		serviceprincipalsecret.SetupGated,
		sqlpermissions.SetupGated,
		token.SetupGated,
		user.SetupGated,
		userinstanceprofile.SetupGated,
		userrole.SetupGated,
		workspaceiamdirectgroupmemberv2.SetupGated,
		workspaceiamgroupv2.SetupGated,
		workspaceiamserviceprincipalv2.SetupGated,
		workspaceiamuserv2.SetupGated,
		workspaceiamworkspaceassignmentv2.SetupGated,
		workspaceiamworkspaceidentitydetailv2.SetupGated,
		endpoint.SetupGated,
		modelserving.SetupGated,
		modelservingprovisionedthroughput.SetupGated,
		onlinestore.SetupGated,
		accountnetworkpolicy.SetupGated,
		accountsettinguserpreferencev2.SetupGated,
		accountsettingv2.SetupGated,
		aibidashboardembeddingaccesspolicysetting.SetupGated,
		aibidashboardembeddingapproveddomainssetting.SetupGated,
		automaticclusterupdateworkspacesetting.SetupGated,
		compliancesecurityprofileworkspacesetting.SetupGated,
		defaultnamespacesetting.SetupGated,
		disablelegacyaccesssetting.SetupGated,
		disablelegacydbfssetting.SetupGated,
		disablelegacyfeaturessetting.SetupGated,
		enhancedsecuritymonitoringworkspacesetting.SetupGated,
		restrictworkspaceadminssetting.SetupGated,
		workspacenetworkoption.SetupGated,
		workspacesettingv2.SetupGated,
		provider.SetupGated,
		recipient.SetupGated,
		share.SetupGated,
		alert.SetupGated,
		alertv2.SetupGated,
		dashboard.SetupGated,
		query.SetupGated,
		sqlalert.SetupGated,
		sqldashboard.SetupGated,
		sqlendpoint.SetupGated,
		sqlglobalconfig.SetupGated,
		sqlquery.SetupGated,
		sqlvisualization.SetupGated,
		sqlwidget.SetupGated,
		warehousesdefaultwarehouseoverride.SetupGated,
		awss3mount.SetupGated,
		azureadlsgen1mount.SetupGated,
		azureadlsgen2mount.SetupGated,
		azureblobmount.SetupGated,
		dbfsfile.SetupGated,
		file.SetupGated,
		mount.SetupGated,
		tagpolicy.SetupGated,
		workspaceentitytagassignment.SetupGated,
		artifactallowlist.SetupGated,
		catalog.SetupGated,
		catalogworkspacebinding.SetupGated,
		connection.SetupGated,
		credential.SetupGated,
		dataqualitymonitor.SetupGated,
		dataqualityrefresh.SetupGated,
		domain.SetupGated,
		entitytagassignment.SetupGated,
		externallocation.SetupGated,
		externalmetadata.SetupGated,
		grant.SetupGated,
		grantmap.SetupGated,
		lakehousemonitor.SetupGated,
		metastore.SetupGated,
		metastoreassignment.SetupGated,
		metastoredataaccess.SetupGated,
		onlinetable.SetupGated,
		policyinfo.SetupGated,
		qualitymonitor.SetupGated,
		qualitymonitorv2.SetupGated,
		registeredmodel.SetupGated,
		rfaaccessrequestdestinations.SetupGated,
		schema.SetupGated,
		secretuc.SetupGated,
		sqltable.SetupGated,
		storagecredential.SetupGated,
		systemschema.SetupGated,
		table.SetupGated,
		volume.SetupGated,
		workspacebinding.SetupGated,
		directory.SetupGated,
		gitcredential.SetupGated,
		globalinitscript.SetupGated,
		notebook.SetupGated,
		notificationdestination.SetupGated,
		repo.SetupGated,
		workspaceconf.SetupGated,
		workspacefile.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		aigatewaymcpservice.SetupWebhookWithManager,
		aigatewaymodelproviderservice.SetupWebhookWithManager,
		aigatewaymodelservice.SetupWebhookWithManager,
		aisearchendpoint.SetupWebhookWithManager,
		aisearchindex.SetupWebhookWithManager,
		knowledgeassistant.SetupWebhookWithManager,
		knowledgeassistantknowledgesource.SetupWebhookWithManager,
		supervisoragent.SetupWebhookWithManager,
		supervisoragenttool.SetupWebhookWithManager,
		app.SetupWebhookWithManager,
		appspace.SetupWebhookWithManager,
		appssettingscustomtemplate.SetupWebhookWithManager,
		customappintegration.SetupWebhookWithManager,
		budgetpolicy.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		clusterpolicy.SetupWebhookWithManager,
		instancepool.SetupWebhookWithManager,
		job.SetupWebhookWithManager,
		library.SetupWebhookWithManager,
		pipeline.SetupWebhookWithManager,
		sandbox.SetupWebhookWithManager,
		databasedatabasecatalog.SetupWebhookWithManager,
		databaseinstance.SetupWebhookWithManager,
		databasesynceddatabasetable.SetupWebhookWithManager,
		instanceprofile.SetupWebhookWithManager,
		mwscredentials.SetupWebhookWithManager,
		mwscustomermanagedkeys.SetupWebhookWithManager,
		mwsnccbinding.SetupWebhookWithManager,
		mwsnccprivateendpointrule.SetupWebhookWithManager,
		mwsnetworkconnectivityconfig.SetupWebhookWithManager,
		mwsnetworks.SetupWebhookWithManager,
		mwsprivateaccesssettings.SetupWebhookWithManager,
		mwsstorageconfigurations.SetupWebhookWithManager,
		mwsvpcendpoint.SetupWebhookWithManager,
		mwsworkspaces.SetupWebhookWithManager,
		disasterrecoveryfailovergroup.SetupWebhookWithManager,
		disasterrecoverystableurl.SetupWebhookWithManager,
		environmentsdefaultworkspacebaseenvironment.SetupWebhookWithManager,
		environmentsworkspacebaseenvironment.SetupWebhookWithManager,
		budget.SetupWebhookWithManager,
		dataclassificationcatalogconfig.SetupWebhookWithManager,
		mwslogdelivery.SetupWebhookWithManager,
		featureengineeringfeature.SetupWebhookWithManager,
		featureengineeringkafkaconfig.SetupWebhookWithManager,
		featureengineeringmaterializedfeature.SetupWebhookWithManager,
		materializedfeaturesfeaturetag.SetupWebhookWithManager,
		mlflowexperiment.SetupWebhookWithManager,
		mlflowmodel.SetupWebhookWithManager,
		mlflowwebhook.SetupWebhookWithManager,
		vectorsearchendpoint.SetupWebhookWithManager,
		vectorsearchindex.SetupWebhookWithManager,
		accountfederationpolicy.SetupWebhookWithManager,
		serviceprincipalfederationpolicy.SetupWebhookWithManager,
		postgresbranch.SetupWebhookWithManager,
		postgrescatalog.SetupWebhookWithManager,
		postgrescdfconfig.SetupWebhookWithManager,
		postgresdataapi.SetupWebhookWithManager,
		postgresdatabase.SetupWebhookWithManager,
		postgresendpoint.SetupWebhookWithManager,
		postgresproject.SetupWebhookWithManager,
		postgresrole.SetupWebhookWithManager,
		postgressnapshotschedule.SetupWebhookWithManager,
		postgressyncedtable.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		accesscontrolruleset.SetupWebhookWithManager,
		accountiamdirectgroupmemberv2.SetupWebhookWithManager,
		accountiamgroupv2.SetupWebhookWithManager,
		accountiamserviceprincipalv2.SetupWebhookWithManager,
		accountiamuserv2.SetupWebhookWithManager,
		accountiamworkspaceassignmentv2.SetupWebhookWithManager,
		entitlements.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		groupinstanceprofile.SetupWebhookWithManager,
		groupmember.SetupWebhookWithManager,
		grouprole.SetupWebhookWithManager,
		ipaccesslist.SetupWebhookWithManager,
		mwspermissionassignment.SetupWebhookWithManager,
		obotoken.SetupWebhookWithManager,
		permissionassignment.SetupWebhookWithManager,
		permissions.SetupWebhookWithManager,
		secret.SetupWebhookWithManager,
		secretacl.SetupWebhookWithManager,
		secretscope.SetupWebhookWithManager,
		serviceprincipal.SetupWebhookWithManager,
		serviceprincipalrole.SetupWebhookWithManager,
		serviceprincipalsecret.SetupWebhookWithManager,
		sqlpermissions.SetupWebhookWithManager,
		token.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		userinstanceprofile.SetupWebhookWithManager,
		userrole.SetupWebhookWithManager,
		workspaceiamdirectgroupmemberv2.SetupWebhookWithManager,
		workspaceiamgroupv2.SetupWebhookWithManager,
		workspaceiamserviceprincipalv2.SetupWebhookWithManager,
		workspaceiamuserv2.SetupWebhookWithManager,
		workspaceiamworkspaceassignmentv2.SetupWebhookWithManager,
		workspaceiamworkspaceidentitydetailv2.SetupWebhookWithManager,
		endpoint.SetupWebhookWithManager,
		modelserving.SetupWebhookWithManager,
		modelservingprovisionedthroughput.SetupWebhookWithManager,
		onlinestore.SetupWebhookWithManager,
		accountnetworkpolicy.SetupWebhookWithManager,
		accountsettinguserpreferencev2.SetupWebhookWithManager,
		accountsettingv2.SetupWebhookWithManager,
		aibidashboardembeddingaccesspolicysetting.SetupWebhookWithManager,
		aibidashboardembeddingapproveddomainssetting.SetupWebhookWithManager,
		automaticclusterupdateworkspacesetting.SetupWebhookWithManager,
		compliancesecurityprofileworkspacesetting.SetupWebhookWithManager,
		defaultnamespacesetting.SetupWebhookWithManager,
		disablelegacyaccesssetting.SetupWebhookWithManager,
		disablelegacydbfssetting.SetupWebhookWithManager,
		disablelegacyfeaturessetting.SetupWebhookWithManager,
		enhancedsecuritymonitoringworkspacesetting.SetupWebhookWithManager,
		restrictworkspaceadminssetting.SetupWebhookWithManager,
		workspacenetworkoption.SetupWebhookWithManager,
		workspacesettingv2.SetupWebhookWithManager,
		provider.SetupWebhookWithManager,
		recipient.SetupWebhookWithManager,
		share.SetupWebhookWithManager,
		alert.SetupWebhookWithManager,
		alertv2.SetupWebhookWithManager,
		dashboard.SetupWebhookWithManager,
		query.SetupWebhookWithManager,
		sqlalert.SetupWebhookWithManager,
		sqldashboard.SetupWebhookWithManager,
		sqlendpoint.SetupWebhookWithManager,
		sqlglobalconfig.SetupWebhookWithManager,
		sqlquery.SetupWebhookWithManager,
		sqlvisualization.SetupWebhookWithManager,
		sqlwidget.SetupWebhookWithManager,
		warehousesdefaultwarehouseoverride.SetupWebhookWithManager,
		awss3mount.SetupWebhookWithManager,
		azureadlsgen1mount.SetupWebhookWithManager,
		azureadlsgen2mount.SetupWebhookWithManager,
		azureblobmount.SetupWebhookWithManager,
		dbfsfile.SetupWebhookWithManager,
		file.SetupWebhookWithManager,
		mount.SetupWebhookWithManager,
		tagpolicy.SetupWebhookWithManager,
		workspaceentitytagassignment.SetupWebhookWithManager,
		artifactallowlist.SetupWebhookWithManager,
		catalog.SetupWebhookWithManager,
		catalogworkspacebinding.SetupWebhookWithManager,
		connection.SetupWebhookWithManager,
		credential.SetupWebhookWithManager,
		dataqualitymonitor.SetupWebhookWithManager,
		dataqualityrefresh.SetupWebhookWithManager,
		domain.SetupWebhookWithManager,
		entitytagassignment.SetupWebhookWithManager,
		externallocation.SetupWebhookWithManager,
		externalmetadata.SetupWebhookWithManager,
		grant.SetupWebhookWithManager,
		grantmap.SetupWebhookWithManager,
		lakehousemonitor.SetupWebhookWithManager,
		metastore.SetupWebhookWithManager,
		metastoreassignment.SetupWebhookWithManager,
		metastoredataaccess.SetupWebhookWithManager,
		onlinetable.SetupWebhookWithManager,
		policyinfo.SetupWebhookWithManager,
		qualitymonitor.SetupWebhookWithManager,
		qualitymonitorv2.SetupWebhookWithManager,
		registeredmodel.SetupWebhookWithManager,
		rfaaccessrequestdestinations.SetupWebhookWithManager,
		schema.SetupWebhookWithManager,
		secretuc.SetupWebhookWithManager,
		sqltable.SetupWebhookWithManager,
		storagecredential.SetupWebhookWithManager,
		systemschema.SetupWebhookWithManager,
		table.SetupWebhookWithManager,
		volume.SetupWebhookWithManager,
		workspacebinding.SetupWebhookWithManager,
		directory.SetupWebhookWithManager,
		gitcredential.SetupWebhookWithManager,
		globalinitscript.SetupWebhookWithManager,
		notebook.SetupWebhookWithManager,
		notificationdestination.SetupWebhookWithManager,
		repo.SetupWebhookWithManager,
		workspaceconf.SetupWebhookWithManager,
		workspacefile.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
