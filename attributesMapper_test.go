package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttMap_Ok(t *testing.T) {
	m := defaultAttributesMapper{}
	source := `<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE ObjectMetadata SYSTEM \"/SysConfig/Classify/FTStories/classify.dtd\"><ObjectMetadata>\n    <EditorialDisplayIndexing>\n        <DILeadCompanies/>\n        <DITemporaryCompanies>\n            <DITemporaryCompany>\n                <DICoTempCode/>\n                <DICoTempDescriptor/>\n                <DICoTickerCode/>\n            </DITemporaryCompany>\n        </DITemporaryCompanies>\n        <DIFTSEGlobalClassifications/>\n        <DIStockExchangeIndices/>\n        <DIHotTopics/>\n        <DIHeadlineCopy>Barclays shares climb as Exane sees good news in history</DIHeadlineCopy>\n        <DIBylineCopy>Bryce Elder</DIBylineCopy>\n        <DIFTNPSections/>\n        \n        \n    <DIFirstParCopy>UK stocks slip to their worst weekly performance since mid-January</DIFirstParCopy><DIMasterImgFileRef>/FT/Graphics/Online/Z_Undefined/2017/03/MAS_Companies_BRIEF_Barclays_5079C30340AE4BD5%208729065EDCDB8156_20170228181405.jpg?uuid=13eada38-1e01-11e7-afad-3fffe3198a53</DIMasterImgFileRef></EditorialDisplayIndexing>\n    <OutputChannels>\n        <DIFTN>\n            <DIFTNPublicationDate/>\n            <DIFTNZoneEdition/>\n            <DIFTNPage/>\n            <DIFTNTimeEdition/>\n            <DIFTNFronts/>\n        </DIFTN>\n        <DIFTcom>\n            <DIFTcomWebType>story</DIFTcomWebType>\n            <DIFTcomDisplayCodes>\n                <DIFTcomDisplayCodeRank1/>\n                <DIFTcomDisplayCodeRank2/>\n            </DIFTcomDisplayCodes>\n            <DIFTcomSubscriptionLevel>0</DIFTcomSubscriptionLevel>\n            <DIFTcomUpdateTimeStamp>False</DIFTcomUpdateTimeStamp>\n            <DIFTcomIndexAndSynd>true</DIFTcomIndexAndSynd>\n            <DIFTcomSafeToSyndicate>True</DIFTcomSafeToSyndicate>\n            <DIFTcomInitialPublication>20170401152914</DIFTcomInitialPublication>\n            <DIFTcomLastPublication>20170410152914</DIFTcomLastPublication>\n            <DIFTcomSuppresInlineAds>False</DIFTcomSuppresInlineAds>\n            <DIFTcomMap>True</DIFTcomMap>\n            <DIFTcomDisplayStyle>Normal</DIFTcomDisplayStyle>\n            <DIFTcomFeatureType>Normal</DIFTcomFeatureType>\n            <DIFTcomMarkDeleted>False</DIFTcomMarkDeleted>\n            <DIFTcomMakeUnlinkable>False</DIFTcomMakeUnlinkable>\n            <isBestStory>0</isBestStory>\n\t\t\t<isContentPackage>False</isContentPackage>\n            <editorsPick>No</editorsPick>\n            <scoop>No</scoop>\n            <exclusive>No</exclusive>\n            <breakingNews>No</breakingNews>\n            <DIFTcomCMRId/>\n            <DIFTcomCMRHint/>\n            <DIFTcomCMR>\n                <DIFTcomCMRPrimarySection/>\n                <DIFTcomCMRPrimarySectionId/>\n                <DIFTcomCMRPrimaryTheme/>\n                <DIFTcomCMRPrimaryThemeId/>\n                <DIFTcomCMRBrand/>\n                <DIFTcomCMRBrandId/>\n                <DIFTcomCMRGenre/>\n                <DIFTcomCMRGenreId/>\n                <DIFTcomCMRMediaType/>\n                <DIFTcomCMRMediaTypeId/>\n            </DIFTcomCMR>\n            \n            \n            \n            \n            \n            \n        <DIFTcomECPositionInText>Default</DIFTcomECPositionInText><DIFTcomHideECLevel1>False</DIFTcomHideECLevel1><DIFTcomHideECLevel2>False</DIFTcomHideECLevel2><DIFTcomHideECLevel3>True</DIFTcomHideECLevel3><DIFTcomDiscussion>True</DIFTcomDiscussion><DIFTcomArticleImage>Article size</DIFTcomArticleImage></DIFTcom>\n        <DISyndication>\n            <DISyndBeenCopied>False</DISyndBeenCopied>\n            <DISyndEdition>USA</DISyndEdition>\n            <DISyndStar>01</DISyndStar>\n            <DISyndChannel/>\n            <DISyndArea/>\n            <DISyndCategory/>\n        </DISyndication>\n    </OutputChannels>\n    <EditorialNotes>\n        <Language>English</Language>\n        <Author>warrenj_new</Author>\n        <Guides/>\n        <Editor/>\n        <Sources>\n            <Source title=\"Financial Times\">\n                <SourceCode>FT</SourceCode>\n                <SourceDescriptor>Financial Times</SourceDescriptor>\n                <SourceOnlineInclusion>True</SourceOnlineInclusion>\n                <SourceCanBeSyndicated>True</SourceCanBeSyndicated>\n            </Source>\n        </Sources>\n        <WordCount>461</WordCount>\n        <CreationDate>20170410141951</CreationDate>\n        <EmbargoDate/>\n        <ExpiryDate>20170410141951</ExpiryDate>\n        <ObjectLocation>/FT/Content/World News/Stories/Live/testimageset.xml</ObjectLocation>\n        <OriginatingStory>c16e8abe-1df8-11e7-942c-4a4c42b3072d</OriginatingStory>\n        \n    <CCMS><CCMSCommissionRefNo/><CCMSContributorRefNo/><CCMSContributorFullName/><CCMSContributorInclude/><CCMSContributorRights>4</CCMSContributorRights><CCMSFilingDate/><CCMSProposedPublishingDate/></CCMS></EditorialNotes>\n    <WiresIndexing>\n        <category/>\n        <Keyword/>\n        <char_count/>\n        <priority/>\n        <basket/>\n        <title/>\n        <Version/>\n        <story_num/>\n        <file_name/>\n        <serviceid/>\n        <entry_date/>\n        <ref_field/>\n        <take_num/>\n    </WiresIndexing>\n    <DataFactoryIndexing>\n        <ADRIS_MetaData>\n            <IndexSuccess/>\n            <StartTime/>\n            <EndTime/>\n        </ADRIS_MetaData>\n        <DFMajorCompanies/>\n        <DFMinorCompanies/>\n        <DFNAICS/>\n        <DFWPMIndustries/>\n        <DFFTSEGlobalClassifications/>\n        <DFStockExchangeIndices/>\n        <DFSubjects/>\n        <DFCountries/>\n        <DFRegions/>\n        <DFWebRegions/>\n        <DFWPMRegions/>\n        <DFProvinces/>\n        <DFFTcomDisplayCodes/>\n        <DFFTSections/>\n    </DataFactoryIndexing>\n</ObjectMetadata>`
	actualAttributes, err := m.Map(source)
	assert.NoError(t, err, "Error wasn't expected during mapping")
	expectedOutputChannels := OutputChannels{
		DIFTcom{
			DIFTcomLastPublication:    "20170410152914",
			DIFTcomInitialPublication: "20170401152914",
		},
	}
	assert.Equal(t, expectedOutputChannels, actualAttributes.OutputChannels)
}

func TestAttMap_NOk(t *testing.T) {
	m := defaultAttributesMapper{}
	_, err := m.Map("<doc><stor**ERROR**</story></doc>")
	assert.Error(t, err, "Error was expected during unmarshal")
}
