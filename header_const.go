package rpm

type SignatureTag int32
type HeaderTag int32

// HEADER_IMAGE 		61
// HEADER_SIGNATURES 	62
// HEADER_IMMUTABLE 	63
// HEADER_REGIONS 		64
// HEADER_I18NTABLE 	100
// HEADER_SIGBASE 		256
// HEADER_TAGBASE 		1000

const (
	HeaderTagNotFound         HeaderTag = -1
	HeaderTagHeaderImage      HeaderTag = 61
	HeaderTagHeaderSignatures HeaderTag = 62
	HeaderTagHeaderImmutable  HeaderTag = 63
	HeaderTagHeaderRegions    HeaderTag = 64
	HeaderTagHeaderI18nTable  HeaderTag = 100
)

const (
	HeaderTagSigBase HeaderTag = iota + 256
	HeaderTagSigSize
	HeaderTagSigLEMD5_1
	HeaderTagSigPGP
	HeaderTagSigLEMD5_2
	HeaderTagSigMD5
	HeaderTagSigGPG
	HeaderTagSigPGP5
	HeaderTagSigBadSHA1_1
	HeaderTagSigBadSHA1_2
	HeaderTagPubKeys
	HeaderTagDSAHeader
	HeaderTagRSAHeader
	HeaderTagSHA1Header
	HeaderTagLongSigSize
	HeaderTagLongArchiveSize

	HeaderTagPkgID    = HeaderTagSigMD5
	HeaderTagHeaderID = HeaderTagSHA1Header
)

const (
	HeaderTagName = iota + 1000
	HeaderTagVersion
	HeaderTagRelease
	HeaderTagEpoch

	HeaderTagSummary
	HeaderTagDescription
	HeaderTagBuildTime
	HeaderTagBuildHost
	HeaderTagInstallTime
	HeaderTagSize
	HeaderTagDistribution
	HeaderTagVendor
	HeaderTagGIF
	HeaderTagXPM
	HeaderTagLicense
	HeaderTagPackager
	HeaderTagGroup
	HeaderTagChangelog
	HeaderTagSource
	HeaderTagPatch
	HeaderTagURL
	HeaderTagOS
	HeaderTagArch
	HeaderTagPreInstall
	HeaderTagPostInstall
	HeaderTagPreUninstall
	HeaderTagPostUninstall
	HeaderTagOldFileNames
	HeaderTagFileSizes
	HeaderTagFileStates
	HeaderTagFileModes
	HeaderTagFileUIDs
	HeaderTagFileGIDs
	HeaderTagFileRDevs
	HeaderTagFileMTimes
	HeaderTagFileDigests

	HeaderTagFileLinkTOS
	HeaderTagFileFlags
	HeaderTagRoot
	HeaderTagFileUserName
	HeaderTagFileGroupName
	HeaderTagExclude
	HeaderTagExclusive
	HeaderTagIcon
	HeaderTagSourceRPM
	HeaderTagFileVerifyFlags
	HeaderTagArchiveSize
	HeaderTagProvideName

	HeaderTagRequireFlags
	HeaderTagRequireName

	HeaderTagRequireVersion
	HeaderTagNoSource
	HeaderTagNoPatch
	HeaderTagConflictFlags
	HeaderTagConflictName

	HeaderTagConflictVersion
	HeaderTagDefaultPrefix
	HeaderTagBuildRoot
	HeaderTagInstallPrefix
	HeaderTagExcludeArch
	HeaderTagExcludeOS
	HeaderTagExclusiveArch
	HeaderTagExclusiveOS
	HeaderTagAutoReqProv
	HeaderTagRPMVersion
	HeaderTagTriggerScripts
	HeaderTagTriggerName
	HeaderTagTriggerVersion
	HeaderTagTriggerFlags
	HeaderTagTriggerIndex
	HeaderTagVerifyScript
	HeaderTagChangelogTime
	HeaderTagChangelogName
	HeaderTagChangelogText
	HeaderTagBrokenMD5
	HeaderTagPreReq
	HeaderTagPreInProg
	HeaderTagPostInstallProg
	HeaderTagPreUninstallProg
	HeaderTagPostUninstallProg
	HeaderTagBuildArchs
	HeaderTagObsoleteName

	HeaderTagVerifyScriptProg
	HeaderTagTriggerScriptProg
	HeaderTagDocDir
	HeaderTagCookie
	HeaderTagFileDevices
	HeaderTagFileINodes
	HeaderTagFileLangs
	HeaderTagPrefixes
	HeaderTagInstallPrefixes
	HeaderTagTriggerInstall
	HeaderTagTriggerUninstall
	HeaderTagTriggerPostUninstall
	HeaderTagAutoReq
	HeaderTagAutoProv
	HeaderTagCapability
	HeaderTagSourcePackage
	HeaderTagOldOrigFilenames
	HeaderTagBuildPreReq
	HeaderTagBuildRequires
	HeaderTagBuildConflicts
	HeaderTagBuildMacros
	HeaderTagProvideFlags
	HeaderTagProvideVersion
	HeaderTagObsoleteFlags
	HeaderTagObsoleteVersion
	HeaderTagDirIndexes
	HeaderTagBaseNames
	HeaderTagDirNames
	HeaderTagOrigDirIndexes
	HeaderTagOrigBaseNames
	HeaderTagOrigDirNames
	HeaderTagOptFlags
	HeaderTagDistURL
	HeaderTagPayloadFormat
	HeaderTagPayloadCompressor
	HeaderTagPayloadFlags
	HeaderTagInstallColor
	HeaderTagInstallTID
	HeaderTagRemoveTID
	HeaderTagSHA1RHN
	HeaderTagRHNPlatform
	HeaderTagPatchesName
	HeaderTagPatchesFlags
	HeaderTagPatchesVersion
	HeaderTagCacheCTime
	HeaderTagCachePkgPath
	HeaderTagCachePkgSize
	HeaderTagCachePkgMTime
	HeaderTagFileColors
	HeaderTagFileClass
	HeaderTagClassDict
	HeaderTagFileDependsX
	HeaderTagFileDependsN
	HeaderTagDependsDict
	HeaderTagSourcePkgID
	HeaderTagFileContexts
	HeaderTagFSContexts
	HeaderTagREContexts
	HeaderTagPolicies
	HeaderTagPreTrans
	HeaderTagPostTrans
	HeaderTagPreTransPRog
	HeaderTagPostTransProg
	HeaderTagDistTag
	HeaderTagOldSuggestsName

	HeaderTagOldSuggestsVersion
	HeaderTagOldSuggestsFlags
	HeaderTagOldEnhancesName

	HeaderTagOldEnhancesVersion
	HeaderTagOldEnhancesFlags
	HeaderTagPriority
	HeaderTagCVSID

	HeaderTagBlinkPkgID
	HeaderTagBlinkHdrID
	HeaderTagBlinkNEVRA
	HeaderTagFlinkPkgID
	HeaderTagFlinkHdrID
	HeaderTagFlinkNEVRA
	HeaderTagPackageOrigin
	HeaderTagTriggerPreInstall
	HeaderTagBuildSuggests
	HeaderTagBuildEnhances
	HeaderTagScriptStates
	HeaderTagScriptMetrics
	HeaderTagBuildCPUClock
	HeaderTagFileDigestAlgos
	HeaderTagVariants
	HeaderTagXMajor
	HeaderTagXMinor
	HeaderTagRepoTag
	HeaderTagKeywords
	HeaderTagBuildPlatforms
	HeaderTagPackageColor
	HeaderTagPackagePrefColor
	HeaderTagXAttrsDict
	HeaderTagFileXAttrsX
	HeaderTagDepAttrsDict
	HeaderTagConflictAttrsX
	HeaderTagObsoleteAttrsX
	HeaderTagProvideAttrsX
	HeaderTagRequireAttrsX
	HeaderTagBuildProvides
	HeaderTagBuildObsoletes
	HeaderTagDBInstance
	HeaderTagNVRA

	HeaderTagN = HeaderTagName
	HeaderTagV = HeaderTagVersion
	HeaderTagR = HeaderTagRelease
	HeaderTagE = HeaderTagEpoch

	HeaderTagFileMD5s = HeaderTagFileDigests

	HeaderTagProvides = HeaderTagProvideName
	HeaderTagP        = HeaderTagProvideName

	HeaderTagRequires = HeaderTagRequireName

	HeaderTagObsoletes = HeaderTagObsoleteName
	HeaderTagO         = HeaderTagObsoleteName

	HeaderTagOldSuggests = HeaderTagOldSuggestsName

	HeaderTagOldEnhances = HeaderTagOldEnhancesName

	HeaderTagSVNID = HeaderTagCVSID
)
