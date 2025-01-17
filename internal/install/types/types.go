// Code generated by tutone: DO NOT EDIT
package types

// OpenInstallationCategory - Categorization of a Quickstart
type OpenInstallationCategory string

var OpenInstallationCategoryTypes = struct {
	// COMMUNITY category
	COMMUNITY OpenInstallationCategory
	// NEWRELIC category
	NEWRELIC OpenInstallationCategory
	// Verified Observbility Pack
	VERIFIED OpenInstallationCategory
}{
	// COMMUNITY category
	COMMUNITY: "COMMUNITY",
	// NEWRELIC category
	NEWRELIC: "NEWRELIC",
	// Verified Observbility Pack
	VERIFIED: "VERIFIED",
}

// OpenInstallationObservabilityPackLevel - Level categorization of Observability Pack
type OpenInstallationObservabilityPackLevel string

var OpenInstallationObservabilityPackLevelTypes = struct {
	// Community Observability Pack
	COMMUNITY OpenInstallationObservabilityPackLevel
	// New Relic Observability Pack
	NEWRELIC OpenInstallationObservabilityPackLevel
	// Verified Observbility Pack
	VERIFIED OpenInstallationObservabilityPackLevel
}{
	// Community Observability Pack
	COMMUNITY: "COMMUNITY",
	// New Relic Observability Pack
	NEWRELIC: "NEWRELIC",
	// Verified Observbility Pack
	VERIFIED: "VERIFIED",
}

// OpenInstallationOperatingSystem - Operating System of target environment
type OpenInstallationOperatingSystem string

var OpenInstallationOperatingSystemTypes = struct {
	// MacOS operating system
	DARWIN OpenInstallationOperatingSystem
	// Linux-based operating system
	LINUX OpenInstallationOperatingSystem
	// Windows operating system
	WINDOWS OpenInstallationOperatingSystem
}{
	// MacOS operating system
	DARWIN: "DARWIN",
	// Linux-based operating system
	LINUX: "LINUX",
	// Windows operating system
	WINDOWS: "WINDOWS",
}

// OpenInstallationPlatform - Operating System distribution
type OpenInstallationPlatform string

var OpenInstallationPlatformTypes = struct {
	// Amazon Linux operating system
	AMAZON OpenInstallationPlatform
	// CentOS operating system
	CENTOS OpenInstallationPlatform
	// Debian operating system
	DEBIAN OpenInstallationPlatform
	// RedHat Enterprise Linux operating system
	REDHAT OpenInstallationPlatform
	// SUSE operating system
	SUSE OpenInstallationPlatform
	// Ubuntu operating system
	UBUNTU OpenInstallationPlatform
}{
	// Amazon Linux operating system
	AMAZON: "AMAZON",
	// CentOS operating system
	CENTOS: "CENTOS",
	// Debian operating system
	DEBIAN: "DEBIAN",
	// RedHat Enterprise Linux operating system
	REDHAT: "REDHAT",
	// SUSE operating system
	SUSE: "SUSE",
	// Ubuntu operating system
	UBUNTU: "UBUNTU",
}

// OpenInstallationPlatformFamily - Operating System distribution family
type OpenInstallationPlatformFamily string

var OpenInstallationPlatformFamilyTypes = struct {
	// Debian distribution family
	DEBIAN OpenInstallationPlatformFamily
	// RHEL distribution family
	RHEL OpenInstallationPlatformFamily
	// openSUSE distribution family
	SUSE OpenInstallationPlatformFamily
}{
	// Debian distribution family
	DEBIAN: "DEBIAN",
	// RHEL distribution family
	RHEL: "RHEL",
	// openSUSE distribution family
	SUSE: "SUSE",
}

// OpenInstallationStability - Stability level of recipe
type OpenInstallationStability string

var OpenInstallationStabilityTypes = struct {
	// Recipe is disabled
	DISABLED OpenInstallationStability
	// Recipe is experimental
	EXPERIMENTAL OpenInstallationStability
	// Recipe is stable
	STABLE OpenInstallationStability
}{
	// Recipe is disabled
	DISABLED: "DISABLED",
	// Recipe is experimental
	EXPERIMENTAL: "EXPERIMENTAL",
	// Recipe is stable
	STABLE: "STABLE",
}

// OpenInstallationSuccessLinkType - Success link type
type OpenInstallationSuccessLinkType string

var OpenInstallationSuccessLinkTypeTypes = struct {
	// Explorer link
	EXPLORER OpenInstallationSuccessLinkType
	// Host entity link
	HOST OpenInstallationSuccessLinkType
}{
	// Explorer link
	EXPLORER: "EXPLORER",
	// Host entity link
	HOST: "HOST",
}

// OpenInstallationTargetType - Installation target type
type OpenInstallationTargetType string

var OpenInstallationTargetTypeTypes = struct {
	// APM agent installation
	APPLICATION OpenInstallationTargetType
	// Cloud provider installation
	CLOUD OpenInstallationTargetType
	// Docker container installation
	DOCKER OpenInstallationTargetType
	// Bare metal, virtual machine, or host-based installation
	HOST OpenInstallationTargetType
	// Kubernetes installation
	KUBERNETES OpenInstallationTargetType
	// Serverless installation
	SERVERLESS OpenInstallationTargetType
}{
	// APM agent installation
	APPLICATION: "APPLICATION",
	// Cloud provider installation
	CLOUD: "CLOUD",
	// Docker container installation
	DOCKER: "DOCKER",
	// Bare metal, virtual machine, or host-based installation
	HOST: "HOST",
	// Kubernetes installation
	KUBERNETES: "KUBERNETES",
	// Serverless installation
	SERVERLESS: "SERVERLESS",
}

// OpenInstallationAttributes - Custom event data attributes
type OpenInstallationAttributes struct {
	// Built-in parsing rulesets
	Logtype string `json:"logtype,omitempty"`
}

// OpenInstallationDocsStitchedFields - Search for and execute installation of additional instrumentation and integrations
type OpenInstallationDocsStitchedFields struct {
	// Fetch a observabilityPack by ID
	ObservabilityPack OpenInstallationObservabilityPack `json:"observabilityPack,omitempty"`
	// Search all Observability Packs
	ObservabilityPackSearch OpenInstallationObservabilityPackResults `json:"observabilityPackSearch,omitempty"`
	// Fetch a recipe by ID
	Recipe OpenInstallationRecipe `json:"recipe,omitempty"`
	// Search all recipes
	RecipeSearch OpenInstallationRecipeListResult `json:"recipeSearch,omitempty"`
	// Get recommendations on what instrumentation to add
	Recommendations OpenInstallationRecommendationsResult `json:"recommendations,omitempty"`
}

// OpenInstallationInstallTarget - Unique set of attributes which represent an install target
type OpenInstallationInstallTarget struct {
	// OS kernel architecture
	KernelArch string `json:"kernelArch,omitempty"`
	// OS kernel version
	KernelVersion string `json:"kernelVersion,omitempty"`
	// Operating system
	Os OpenInstallationOperatingSystem `json:"os,omitempty"`
	// OS distribution
	Platform OpenInstallationPlatform `json:"platform,omitempty"`
	// OS distribution family
	PlatformFamily OpenInstallationPlatformFamily `json:"platformFamily,omitempty"`
	// OS distribution version
	PlatformVersion string `json:"platformVersion,omitempty"`
	// Target type
	Type OpenInstallationTargetType `json:"type,omitempty"`
}

// OpenInstallationLogMatch - Matches partial list of the Log forwarding parameters
type OpenInstallationLogMatch struct {
	// List of custom attributes, as key-value pairs, that can be used to send additional data with the logs which you can then query.
	Attributes OpenInstallationAttributes `json:"attributes,omitempty"`
	// Path to the log file or files.
	File string `json:"file,omitempty"`
	// Name of the log or logs.
	Name string `json:"name"`
	// Regular expression for filtering records.
	Pattern string `json:"pattern,omitempty"`
	// Service name (Linux Only).
	Systemd string `json:"systemd,omitempty"`
}

// OpenInstallationObservabilityPack - An Observability Pack (collection of related Observability features)
type OpenInstallationObservabilityPack struct {
	// List of contributors for this Observability Pack
	Authors []string `json:"authors"`
	// List of Dashboards contained in this Observability Pack
	Dashboards []OpenInstallationObservabilityPackDashboard `json:"dashboards"`
	// A short form description for this Observability Pack. Used throughout the platform when displaying the pack.
	Description string `json:"description"`
	// A unique Observability Pack identifier
	ID string `json:"id"`
	// URL of icon for this Observability Pack
	IconURL string `json:"iconUrl,omitempty"`
	// Level categorization of Observability Pack
	Level OpenInstallationObservabilityPackLevel `json:"level"`
	// URL of logo for this Observability Pack
	LogoURL string `json:"logoUrl,omitempty"`
	// The name of this Observability Pack
	Name string `json:"name"`
	// URL of website for this Observability Pack
	WebsiteURL string `json:"websiteUrl,omitempty"`
}

// OpenInstallationObservabilityPackDashboard - An Observability Pack Dashboard
type OpenInstallationObservabilityPackDashboard struct {
	// Description of this Dashboard
	Description string `json:"description"`
	// Name of this Dashboard
	Name string `json:"name"`
	// List of screenshot URLs related to this Dashboard
	Screenshots []string `json:"screenshots"`
	// URL of Dashboard JSON definition
	URL string `json:"url"`
}

// OpenInstallationObservabilityPackFilter - Metadata used to filter for Observability Packs
type OpenInstallationObservabilityPackFilter struct {
	// Level categorization of Observability Pack
	Level OpenInstallationObservabilityPackLevel `json:"level,omitempty"`
	// Name of the Observability Pack
	Name string `json:"name,omitempty"`
}

// OpenInstallationObservabilityPackInputCriteria - Input for searching installable Observability Packs
type OpenInstallationObservabilityPackInputCriteria struct {
	// Support level of an Observability Pack
	Level OpenInstallationObservabilityPackLevel `json:"level,omitempty"`
	// The name of the Observability Pack
	Name string `json:"name,omitempty"`
}

// OpenInstallationObservabilityPackResults - A data structure that contains the detailed response of an Observability Pack search.
//
// The direct search result is available through results. Information about the query itself is available through count.
type OpenInstallationObservabilityPackResults struct {
	// Number of results
	Count int `json:"count,omitempty"`
	// Cursor pointing to the end of the current page of Observability Packs. Null if final page.
	NextCursor string `json:"nextCursor,omitempty"`
	// The paginated results of the Observability Pack search
	Results OpenInstallationObservabilityPackSearchResult `json:"results,omitempty"`
}

// OpenInstallationObservabilityPackSearchResult - A data structure that contains Observability Packs and the related components.
type OpenInstallationObservabilityPackSearchResult struct {
	// List of Observability Packs returned for this search
	ObservabilityPacks []OpenInstallationObservabilityPack `json:"observabilityPacks"`
}

// OpenInstallationPostInstallConfiguration - Optional post-install configuration items
type OpenInstallationPostInstallConfiguration struct {
	// Message/Docs notice displayed to user after running the recipe
	Info string `json:"info,omitempty"`
}

// OpenInstallationPreInstallConfiguration - Optional pre-install configuration items
type OpenInstallationPreInstallConfiguration struct {
	// Message/Docs notice displayed to user prior to running recipe
	Info string `json:"info,omitempty"`
	// Message/Docs notice displayed to user prior to running recipe
	Prompt string `json:"prompt,omitempty"`
	// Script block to be executed during system discovery, a successful exit status will mark the recipe for execution
	RequireAtDiscovery string `json:"requireAtDiscovery,omitempty"`
}

// OpenInstallationProcessDetailInput - Process details
type OpenInstallationProcessDetailInput struct {
	// Process name
	Name string `json:"name,omitempty"`
}

// OpenInstallationQuickstartEntityType - Entity type relation for Quickstart
type OpenInstallationQuickstartEntityType struct {
	// Domain of Entity
	Domain string `json:"domain"`
	// Type of Entity
	Type string `json:"type"`
}

// OpenInstallationQuickstartsFilter - Metadata used to filter for Quickstarts
type OpenInstallationQuickstartsFilter struct {
	// Categorization of Quickstart
	Category OpenInstallationCategory `json:"category,omitempty"`
	// Entity relationship for Quickstart
	EntityType OpenInstallationQuickstartEntityType `json:"entityType,omitempty"`
	// Name of the Quickstart
	Name string `json:"name"`
}

// OpenInstallationRecipe - Installation instructions and definition of an instrumentation integration
type OpenInstallationRecipe struct {
	// Named list of dependencies for this recipe
	Dependencies []string `json:"dependencies"`
	// Description of the recipe
	Description string `json:"description"`
	// Friendly name of the integration
	DisplayName string `json:"displayName,omitempty"`
	// The full contents of the recipe file (yaml)
	File string `json:"file"`
	// The ID
	ID string `json:"id,omitempty"`
	// List of variables to prompt for input from the user
	InputVars []OpenInstallationRecipeInputVariable `json:"inputVars"`
	// Go-task's taskfile definition (see https://taskfile.dev/#/usage)
	Install string `json:"install"`
	// Object representing the intended install target
	InstallTargets []OpenInstallationRecipeInstallTarget `json:"installTargets"`
	// Tags
	Keywords []string `json:"keywords"`
	// # Partial list of possible Log forwarding parameters
	LogMatch []OpenInstallationLogMatch `json:"logMatch"`
	// Short unique handle for the name of the integration
	Name string `json:"name,omitempty"`
	// Metadata used to recommend and install Observability Packs
	ObservabilityPacks []OpenInstallationObservabilityPackFilter `json:"observabilityPacks"`
	// Object representing optional post-install configuration items
	PostInstall OpenInstallationPostInstallConfiguration `json:"postInstall,omitempty"`
	// Object representing optional pre-install configuration items
	PreInstall OpenInstallationPreInstallConfiguration `json:"preInstall,omitempty"`
	// List of process definitions used to match CLI process detection
	ProcessMatch []string `json:"processMatch"`
	// Metadata used to recommend and install Quickstarts
	Quickstarts OpenInstallationQuickstartsFilter `json:"quickstarts,omitempty"`
	// Github repository url
	Repository string `json:"repository"`
	// Indicates stability level of recipe
	Stability OpenInstallationStability `json:"stability,omitempty"`
	// Metadata to support generating a URL after installation success
	SuccessLinkConfig OpenInstallationSuccessLinkConfig `json:"successLinkConfig,omitempty"`
	// NRQL the newrelic-cli uses to validate this recipe
	// is successfully sending data to New Relic
	ValidationNRQL NRQL `json:"validationNrql,omitempty"`
	// validation url to validate with infra health endpoint
	ValidationURL string `json:"validationUrl,omitempty"`
}

// OpenInstallationRecipeInputVariable - Recipe input variable prompts displayed to the user prior to execution
type OpenInstallationRecipeInputVariable struct {
	// Default value of variable
	Default string `json:"default,omitempty"`
	// Name of the variable
	Name string `json:"name"`
	// Message to present to the user
	Prompt string `json:"prompt,omitempty"`
	// Indicates a password field
	Secret bool `json:"secret,omitempty"`
}

// OpenInstallationRecipeInstallTarget - Matrix of supported installation criteria for this recipe
type OpenInstallationRecipeInstallTarget struct {
	// OS kernel architecture
	KernelArch string `json:"kernelArch,omitempty"`
	// OS kernel version
	KernelVersion string `json:"kernelVersion,omitempty"`
	// Operating system
	Os OpenInstallationOperatingSystem `json:"os,omitempty"`
	// Operating System distribution
	Platform OpenInstallationPlatform `json:"platform,omitempty"`
	// Operating System distribution family
	PlatformFamily OpenInstallationPlatformFamily `json:"platformFamily,omitempty"`
	// OS distribution version
	PlatformVersion string `json:"platformVersion,omitempty"`
	// Target type
	Type OpenInstallationTargetType `json:"type,omitempty"`
}

// OpenInstallationRecipeListResult - List of recipes
type OpenInstallationRecipeListResult struct {
	// Number of recipes returned
	Count int `json:"count,omitempty"`
	// List of recipes
	Results []OpenInstallationRecipe `json:"results,omitempty"`
}

// OpenInstallationRecipeSearchCriteria - Input for searching installable integration recipes
type OpenInstallationRecipeSearchCriteria struct {
	// Friendly name of the integration
	DisplayName string `json:"displayName,omitempty"`
	// Object representing the intended install target
	InstallTarget OpenInstallationInstallTarget `json:"installTarget,omitempty"`
	// Short unique handle for the name of the integration
	Name string `json:"name,omitempty"`
}

// OpenInstallationRecommendationsInput - Criteria for controlling recommendations
type OpenInstallationRecommendationsInput struct {
	// Unique set of attributes which represent an install target
	InstallTarget OpenInstallationInstallTarget `json:"installTarget,omitempty"`
	// List of processes
	ProcessDetails []OpenInstallationProcessDetailInput `json:"processDetails,omitempty"`
}

// OpenInstallationRecommendationsResult - List of instrumentation recommendations
type OpenInstallationRecommendationsResult struct {
	// Number of recipes returned
	Count int `json:"count,omitempty"`
	// List of matching recipes
	Results []OpenInstallationRecipe `json:"results"`
}

// OpenInstallationSuccessLinkConfig - Metadata to support generating a URL after installation success
type OpenInstallationSuccessLinkConfig struct {
	// An optional filter for appending to the URL
	Filter string `json:"filter,omitempty"`
	// The type of the link to generate
	Type OpenInstallationSuccessLinkType `json:"type"`
}

// NRQL - This scalar represents a NRQL query string.
//
// See the [NRQL Docs](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/nrql-resources/nrql-syntax-components-functions) for more information about NRQL syntax.
type NRQL string
