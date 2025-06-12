package database

import (
	"server/models"

	"gorm.io/gorm"
)

type Seeder interface {
	Seed(db *gorm.DB) error
	Clear(db *gorm.DB) error
}

// Lifecycle

type LifecycleSeeder struct{}

func (s LifecycleSeeder) Seed(db *gorm.DB) error {
	lifecycles := []models.Lifecycle{
		{
			Title:       "ELSA Recommendation Platform for Data Science Projects",
			Description: "prototype",
			General: `
## Welcome to the ELSA Recommendation Platform for Data Science Projects

As data science continues to shape the world around us, from healthcare and education to urban planning and policymaking, there is a growing need to ensure that the work we do is not only technically sound, but also ethically responsible, legally compliant, and socially aware.
			
That’s where this platform comes in.
			
The ELSA (Ethical, Legal, and Societal Aspects) Recommendation Platform is designed to support data scientists throughout the entire lifecycle of a project. Whether you're just starting to define your problem, collecting and preparing data, building models, or deploying results in real-world contexts, this platform helps you reflect on and act upon important questions like:

- **Are we using data responsibly?**
- **Have we considered the impact on individuals and communities?**
- **What are the legal requirements or potential risks?**
- **Who might be affected, and how are their interests being considered?**

Instead of asking you to become an expert in ethics, law, or the social sciences, the platform offers practical guidance, curated tools, checklists, and methods tailored to each stage of the data science workflow. It also allows you to document your decisions and reflections in a structured, project-specific journal—helping you keep track of how and when these considerations have been addressed.

Whether you’re part of a research team, a public-sector project, or an industry application, this platform aims to make it easier to build data-driven solutions that are not only innovative but also responsible and trustworthy.

### What does this look like in practice?

Imagine you're building a machine learning model to predict which patients are at risk of developing a chronic illness. In the early planning stages, the platform might prompt you to think about potential biases in your training data, or whether patients have consented to their data being used in this way. You might be directed to a checklist for assessing data protection compliance, or a tool for mapping potential societal impacts.

Later, when you're training and evaluating your model, the platform could recommend fairness metrics or guidelines for documenting model performance across different population groups. As you prepare to deploy the model in a clinical setting, you might receive suggestions on transparency practices or legal frameworks relevant to medical AI tools.

At each step, your reflections and decisions are captured in your project journal—building a transparent record of how ethical, legal, and societal aspects have been considered and addressed.`,
			Introduction: `
## Why Keep an ELSA Journal?

### An ELSA journal is like version control for your ethical reasoning.

It adds structure and traceability to a process that’s otherwise invisible — enabling smarter decisions, easier collaboration, and more robust, responsible systems.

### An ELSA journal supports rigorous, transparent workflows

Just like a lab notebook in experimental science, an ELSA journal captures the decision-making process, including what was considered, what was chosen, and why.

This supports:

- **Reproducibility:** Others (or your future self) can trace ethical and societal reasoning.
- **Transparency:** Your rationale for handling fairness, privacy, or bias is documented, not buried in memory or Slack threads.

### An ELSA journal helps to simplify compliance requirements and reviews

Many data projects now fall under regulatory scrutiny (e.g. GDPR, AI Act, public procurement policies).

The journal:

- Acts as an audit trail for legal or ethical review.
- Helps justify decisions made under uncertainty (e.g. why a certain dataset was excluded, or why a particular fairness metric was used).
- Makes it easier to write final impact assessments, ethics summaries, or model documentation.

### An ELSA journal encourages iterative reflection and course-correction

Keeping track of ELSA-related concerns allows teams to:

- Spot inconsistencies or blind spots early in the workflow.
- Reflect on trade-offs (e.g., accuracy vs. fairness, complexity vs. explainability).
- Capture insights that can feed into model or data revisions down the line.

### An ELSA journal can help to improve cross-disciplinary collaboration

For projects involving ethicists, legal experts, social scientists, or impacted stakeholders:

- The journal creates a shared reference point for discussions.
- It helps translate technical choices into accessible justifications for non-technical audiences.

### An ELSA journal contributes to institutional learning and best practices

Over time, a library of journal entries across projects:

1. Becomes a living archive of lessons learned.
2. Reveals recurring ethical bottlenecks or high-risk points in the pipeline.
3. Helps build internal guidelines or even training materials based on real experience.`,
		},
	}
	return db.Create(&lifecycles).Error
}

func (s LifecycleSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM lifecycles").Error
}

// Phase

type PhaseSeeder struct{}

func (s PhaseSeeder) Seed(db *gorm.DB) error {
	phases := []models.Phase{
		{
			Number:      1,
			Title:       "Problem Definition / Project Scoping",
			Description: "prototype",
			LifecycleID: 1,
		},
		{
			Number:      2,
			Title:       "Data Collection / Acquisition",
			Description: "prototype",
			LifecycleID: 1,
		},
	}
	return db.Create(&phases).Error
}

func (s PhaseSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM phases").Error
}

// Tool

type ToolSeeder struct{}

func (s ToolSeeder) Seed(db *gorm.DB) error {
	tools := []models.Tool{
		{
			Title:       "Value Sensitive Design",
			Description: "A method for integrating stakeholder values into the design process.",
			URL:         "https://vsdesign.org/",
			Cover:       "",
			Tags:        "Values,Stakeholders",
			Type:        "Method",
			Form: `{
					"@id": "https://repo.metadatacenter.org/templates/6c0fbbd3-86b3-415f-aba1-9a15a764e5f8",
					"@type": "https://schema.metadatacenter.org/core/Template",
					"@context": {
						"xsd": "http://www.w3.org/2001/XMLSchema#",
						"pav": "http://purl.org/pav/",
						"bibo": "http://purl.org/ontology/bibo/",
						"oslc": "http://open-services.net/ns/core#",
						"schema": "http://schema.org/",
						"schema:name": {
						"@type": "xsd:string"
						},
						"schema:description": {
						"@type": "xsd:string"
						},
						"pav:createdOn": {
						"@type": "xsd:dateTime"
						},
						"pav:createdBy": {
						"@type": "@id"
						},
						"pav:lastUpdatedOn": {
						"@type": "xsd:dateTime"
						},
						"oslc:modifiedBy": {
						"@type": "@id"
						}
					},
					"type": "object",
					"title": "Value sensitive design template schema",
					"description": "Value sensitive design template schema generated by the CEDAR Template Editor 2.7.7",
					"_ui": {
						"order": [
						"Considerations",
						"start-date",
						"members",
						"public"
						],
						"propertyLabels": {
						"Considerations": "Considerations",
						"start-date": "start-date",
						"members": "members",
						"public": "public"
						},
						"propertyDescriptions": {
						"Considerations": "",
						"start-date": "",
						"members": "Help Text",
						"public": ""
						}
					},
					"properties": {
						"@context": {
						"type": "object",
						"properties": {
							"rdfs": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2000/01/rdf-schema#"
							]
							},
							"xsd": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2001/XMLSchema#"
							]
							},
							"pav": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://purl.org/pav/"
							]
							},
							"schema": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://schema.org/"
							]
							},
							"oslc": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://open-services.net/ns/core#"
							]
							},
							"skos": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2004/02/skos/core#"
							]
							},
							"rdfs:label": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:isBasedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"schema:name": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:description": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"pav:derivedFrom": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:createdOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"pav:createdBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:lastUpdatedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"oslc:modifiedBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"skos:notation": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"Considerations": {
							"enum": [
								"https://schema.metadatacenter.org/properties/04001e61-0bff-421f-badc-d5e6965d19dd"
							]
							},
							"start-date": {
							"enum": [
								"https://schema.metadatacenter.org/properties/698d67bb-4af0-4fb5-97a6-9933f1540d8c"
							]
							},
							"members": {
							"enum": [
								"https://schema.metadatacenter.org/properties/563e4001-a3a6-46ee-9d19-0618ea1f85cd"
							]
							},
							"public": {
							"enum": [
								"https://schema.metadatacenter.org/properties/44791356-5c5e-4e60-ba81-6a7c88e70b8f"
							]
							}
						},
						"required": [
							"xsd",
							"pav",
							"schema",
							"oslc",
							"schema:isBasedOn",
							"schema:name",
							"schema:description",
							"pav:createdOn",
							"pav:createdBy",
							"pav:lastUpdatedOn",
							"oslc:modifiedBy"
						],
						"additionalProperties": false
						},
						"@id": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"@type": {
						"oneOf": [
							{
							"type": "string",
							"format": "uri"
							},
							{
							"type": "array",
							"minItems": 1,
							"items": {
								"type": "string",
								"format": "uri"
							},
							"uniqueItems": true
							}
						]
						},
						"schema:isBasedOn": {
						"type": "string",
						"format": "uri"
						},
						"schema:name": {
						"type": "string",
						"minLength": 1
						},
						"schema:description": {
						"type": "string"
						},
						"pav:derivedFrom": {
						"type": "string",
						"format": "uri"
						},
						"pav:createdOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"pav:createdBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"pav:lastUpdatedOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"oslc:modifiedBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"Considerations": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "Considerations field schema",
						"description": "Considerations field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "textfield"
						},
						"_valueConstraints": {
							"requiredValue": true
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "Considerations",
						"schema:description": "",
						"pav:createdOn": "2025-06-10T00:33:40-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-10T00:33:40-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"@id": "https://repo.metadatacenter.org/template-fields/027abd53-61b7-42db-bb00-bf0608a5318b",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"start-date": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "start-date field schema",
						"description": "start-date field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "temporal",
							"temporalGranularity": "second",
							"timezoneEnabled": true,
							"inputTimeFormat": "12h"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"temporalType": "xsd:dateTime"
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "start-date",
						"schema:description": "",
						"pav:createdOn": "2025-06-10T00:33:40-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-10T00:33:40-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Start Date",
						"@id": "https://repo.metadatacenter.org/template-fields/b4dec18d-6a20-4ed6-8aab-7b13761f1761",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"members": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "members field schema",
						"description": "members field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "numeric"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"numberType": "xsd:decimal"
						},
						"properties": {
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							},
							"@type": {
							"type": "string",
							"format": "uri"
							}
						},
						"required": [
							"@value",
							"@type"
						],
						"schema:name": "members",
						"schema:description": "Help Text",
						"pav:createdOn": "2025-06-10T00:33:40-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-10T00:33:40-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Number of members",
						"@id": "https://repo.metadatacenter.org/template-fields/352c91fa-5f97-4a08-b136-d7b70681f417",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"public": {
						"type": "array",
						"minItems": 1,
						"items": {
							"type": "object",
							"@type": "https://schema.metadatacenter.org/core/TemplateField",
							"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
								"@type": "xsd:string"
							},
							"schema:description": {
								"@type": "xsd:string"
							},
							"skos:prefLabel": {
								"@type": "xsd:string"
							},
							"skos:altLabel": {
								"@type": "xsd:string"
							},
							"pav:createdOn": {
								"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
								"@type": "@id"
							},
							"pav:lastUpdatedOn": {
								"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
								"@type": "@id"
							}
							},
							"title": "public field schema",
							"description": "public field schema generated by the CEDAR Template Editor 2.7.7",
							"_ui": {
							"inputType": "checkbox"
							},
							"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": true,
							"literals": [
								{
								"label": "Yes"
								},
								{
								"label": "No"
								}
							]
							},
							"properties": {
							"@type": {
								"oneOf": [
								{
									"type": "string",
									"format": "uri"
								},
								{
									"type": "array",
									"minItems": 1,
									"items": {
									"type": "string",
									"format": "uri"
									},
									"uniqueItems": true
								}
								]
							},
							"@value": {
								"type": [
								"string",
								"null"
								]
							},
							"rdfs:label": {
								"type": [
								"string",
								"null"
								]
							}
							},
							"required": [
							"@value"
							],
							"additionalProperties": false,
							"pav:createdOn": "2025-06-10T00:33:40-07:00",
							"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
							"pav:lastUpdatedOn": "2025-06-10T00:33:40-07:00",
							"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
							"schema:schemaVersion": "1.6.0",
							"schema:name": "public",
							"schema:description": "",
							"skos:prefLabel": "Is Public?",
							"@id": "https://repo.metadatacenter.org/template-fields/5950f1f3-b226-4fe8-9f85-916adbe51e93",
							"$schema": "http://json-schema.org/draft-04/schema#"
						}
						}
					},
					"required": [
						"@context",
						"@id",
						"schema:isBasedOn",
						"schema:name",
						"schema:description",
						"pav:createdOn",
						"pav:createdBy",
						"pav:lastUpdatedOn",
						"oslc:modifiedBy",
						"Considerations",
						"start-date",
						"members",
						"public"
					],
					"schema:name": "Value sensitive design",
					"schema:description": "",
					"pav:createdOn": "2025-06-10T00:33:40-07:00",
					"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"pav:lastUpdatedOn": "2025-06-10T00:33:40-07:00",
					"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"schema:schemaVersion": "1.6.0",
					"additionalProperties": false,
					"pav:version": "0.0.1",
					"bibo:status": "bibo:draft",
					"$schema": "http://json-schema.org/draft-04/schema#"
					}`,
		},
		{
			Title:       "Problem Framing Canvas",
			Description: "A visual tool to help teams critically reflect on how a problem is framed.",
			URL:         "https://realkm.com/wp-content/uploads/2023/05/Problem-Framing-Canvas-Handbook.pdf",
			Cover:       "",
			Tags:        "Problem,Framing",
			Type:        "Mapping, Tool, Canvas",
		},
		{
			Title:       "Stakeholder Mapping Tool",
			Description: "Tool for identifying and analyzing stakeholder relationships.",
			URL:         "https://www.stakeholdermap.com/",
			Cover:       "",
			Tags:        "Stakeholders",
			Type:        "Mapping, Tool",
		},
		{
			Title:       "UNESCO Ethics of AI Recommendation",
			Description: "Normative framework providing ethical guidelines for AI development and use.",
			URL:         "https://unesdoc.unesco.org/ark:/48223/pf0000381137",
			Cover:       "",
			Tags:        "Ethics",
			Type:        "Guidelines, Framework",
		},
		{
			Title:       "Data Ethics Canvas",
			Description: "Canvas for identifying and addressing ethical issues in data projects.",
			URL:         "https://theodi.org/article/data-ethics-canvas/",
			Cover:       "",
			Tags:        "Ethics",
			Type:        "Mapping, Canvas",
		},
		{
			Title:       "GDPR Compliance Checklist",
			Description: "Checklist for ensuring compliance with the General Data Protection Regulation.",
			URL:         "https://gdpr.eu/checklist/",
			Cover:       "",
			Tags:        "",
			Type:        "Checklist",
		},
		{
			Title:       "FAIR Principles",
			Description: "Guidelines for making data Findable, Accessible, Interoperable, and Reusable.",
			URL:         "https://www.go-fair.org/fair-principles/",
			Cover:       "",
			Tags:        "",
			Type:        "Guidelines",
		},
		{
			Title:       "Datasheets for Datasets",
			Description: "Standardized documentation approach for datasets.",
			URL:         "https://arxiv.org/abs/1803.09010",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "Missing Data Bias Tool",
			Description: "Tool to assess potential bias introduced by missing data.",
			URL:         "https://arxiv.org/pdf/2406.16846",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "Fairlearn",
			Description: "Toolkit to assess and mitigate fairness issues in machine learning.",
			URL:         "https://fairlearn.org/",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "Model Cards",
			Description: "Templates for transparent reporting of model characteristics.",
			URL:         "https://modelcards.withgoogle.com/about",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "AI Fairness 360",
			Description: "IBM's toolkit for detecting and mitigating bias in ML models.",
			URL:         "https://research.ibm.com/blog/ai-fairness-360",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "What-If Tool",
			Description: "Interactive visual tool for exploring ML model performance and fairness.",
			URL:         "https://pair-code.github.io/what-if-tool/",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "Algorithm Impact Assessment (Ada Lovelace Institute)",
			Description: "Framework for assessing the impact of algorithms before deployment (healthcare focus)",
			URL:         "https://www.adalovelaceinstitute.org/resource/aia-user-guide/",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "NIST AI Risk Management Framework",
			Description: "Framework for managing AI-related risks across the lifecycle.",
			URL:         "https://www.nist.gov/itl/ai-risk-management-framework",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
		{
			Title:       "What is AI transparency? A comprehensive guide",
			Description: "Guide to improve AI transparency.",
			URL:         "https://www.zendesk.nl/blog/ai-transparency/#",
			Cover:       "",
			Tags:        "",
			Type:        "",
		},
	}
	return db.Create(&tools).Error
}

func (s ToolSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM tools").Error
}

// Reflection

type ReflectionSeeder struct{}

func (s ReflectionSeeder) Seed(db *gorm.DB) error {
	reflections := []models.Reflection{
		{
			Title:       "Problem Definition / Project Scoping",
			Description: "As data science continues to shape the world around us, from healthcare and education to urban planning and policymaking, there is a growing need to ensure that the work we do is not only technically sound, but also ethically responsible, legally compliant, and socially aware.",
			Form: `{
					"@id": "https://repo.metadatacenter.org/templates/1ec51eb6-d07c-4f25-8299-1972b55b42b7",
					"@type": "https://schema.metadatacenter.org/core/Template",
					"@context": {
						"xsd": "http://www.w3.org/2001/XMLSchema#",
						"pav": "http://purl.org/pav/",
						"bibo": "http://purl.org/ontology/bibo/",
						"oslc": "http://open-services.net/ns/core#",
						"schema": "http://schema.org/",
						"schema:name": {
						"@type": "xsd:string"
						},
						"schema:description": {
						"@type": "xsd:string"
						},
						"pav:createdOn": {
						"@type": "xsd:dateTime"
						},
						"pav:createdBy": {
						"@type": "@id"
						},
						"pav:lastUpdatedOn": {
						"@type": "xsd:dateTime"
						},
						"oslc:modifiedBy": {
						"@type": "@id"
						}
					},
					"type": "object",
					"title": "Boolean template schema",
					"description": "Boolean template schema generated by the CEDAR Template Editor 2.7.1",
					"_ui": {
						"order": [
						"problem_definition",
						"benefit_harm",
						"alternative_framing",
						"consult_stakeholders"
						],
						"propertyLabels": {
						"problem_definition": "problem_definition",
						"benefit_harm": "benefit_harm",
						"alternative_framing": "alternative_framing",
						"consult_stakeholders": "consult_stakeholders"
						},
						"propertyDescriptions": {
						"problem_definition": "",
						"benefit_harm": "",
						"alternative_framing": "Help Text",
						"consult_stakeholders": ""
						}
					},
					"properties": {
						"@context": {
						"type": "object",
						"properties": {
							"rdfs": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2000/01/rdf-schema#"
							]
							},
							"xsd": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2001/XMLSchema#"
							]
							},
							"pav": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://purl.org/pav/"
							]
							},
							"schema": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://schema.org/"
							]
							},
							"oslc": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://open-services.net/ns/core#"
							]
							},
							"skos": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2004/02/skos/core#"
							]
							},
							"rdfs:label": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:isBasedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"schema:name": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:description": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"pav:derivedFrom": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:createdOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"pav:createdBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:lastUpdatedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"oslc:modifiedBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"skos:notation": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"problem_definition": {
							"enum": [
								"https://schema.metadatacenter.org/properties/b9503e1b-f332-43ef-b2d8-7aceb651e65b"
							]
							},
							"benefit_harm": {
							"enum": [
								"https://schema.metadatacenter.org/properties/caf76e23-2773-4623-80cc-0991cdb3bec8"
							]
							},
							"alternative_framing": {
							"enum": [
								"https://schema.metadatacenter.org/properties/b2a81610-e5e9-4e21-9351-c8396521e6f9"
							]
							},
							"consult_stakeholders": {
							"enum": [
								"https://schema.metadatacenter.org/properties/b2db8450-ba8b-46d2-acbe-c746033bd6ff"
							]
							}
						},
						"required": [
							"xsd",
							"pav",
							"schema",
							"oslc",
							"schema:isBasedOn",
							"schema:name",
							"schema:description",
							"pav:createdOn",
							"pav:createdBy",
							"pav:lastUpdatedOn",
							"oslc:modifiedBy"
						],
						"additionalProperties": false
						},
						"@id": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"@type": {
						"oneOf": [
							{
							"type": "string",
							"format": "uri"
							},
							{
							"type": "array",
							"minItems": 1,
							"items": {
								"type": "string",
								"format": "uri"
							},
							"uniqueItems": true
							}
						]
						},
						"schema:isBasedOn": {
						"type": "string",
						"format": "uri"
						},
						"schema:name": {
						"type": "string",
						"minLength": 1
						},
						"schema:description": {
						"type": "string"
						},
						"pav:derivedFrom": {
						"type": "string",
						"format": "uri"
						},
						"pav:createdOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"pav:createdBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"pav:lastUpdatedOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"oslc:modifiedBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"problem_definition": {
						"$schema": "http://json-schema.org/draft-04/schema#",
						"@id": "tmp-1747217452591-61489",
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "problem_definition field schema",
						"description": "problem_definition field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "problem_definition",
						"schema:description": "",
						"pav:createdOn": null,
						"pav:createdBy": null,
						"pav:lastUpdatedOn": null,
						"oslc:modifiedBy": null,
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Have you clearly defined the problem and its purpose?"
						},
						"benefit_harm": {
						"$schema": "http://json-schema.org/draft-04/schema#",
						"@id": "tmp-1747217616288-225186",
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "benefit_harm field schema",
						"description": "benefit_harm field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "benefit_harm",
						"schema:description": "",
						"pav:createdOn": null,
						"pav:createdBy": null,
						"pav:lastUpdatedOn": null,
						"oslc:modifiedBy": null,
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Have you identified who will benefit from this project - and who might be harmed?"
						},
						"alternative_framing": {
						"$schema": "http://json-schema.org/draft-04/schema#",
						"@id": "tmp-1747217763709-372607",
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "alternative_framing field schema",
						"description": "alternative_framing field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "alternative_framing",
						"schema:description": "Help Text",
						"pav:createdOn": null,
						"pav:createdBy": null,
						"pav:lastUpdatedOn": null,
						"oslc:modifiedBy": null,
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Have you considered alternative ways to frame the problem?"
						},
						"consult_stakeholders": {
						"$schema": "http://json-schema.org/draft-04/schema#",
						"@id": "tmp-1747217837267-446166",
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "consult_stakeholders field schema",
						"description": "consult_stakeholders field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "consult_stakeholders",
						"schema:description": "",
						"pav:createdOn": null,
						"pav:createdBy": null,
						"pav:lastUpdatedOn": null,
						"oslc:modifiedBy": null,
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Have you engaged or consulted stakeholders, specially those affected?"
						}
					},
					"required": [
						"@context",
						"@id",
						"schema:isBasedOn",
						"schema:name",
						"schema:description",
						"pav:createdOn",
						"pav:createdBy",
						"pav:lastUpdatedOn",
						"oslc:modifiedBy",
						"problem_definition",
						"benefit_harm",
						"alternative_framing",
						"consult_stakeholders"
					],
					"schema:name": "Boolean",
					"schema:description": "",
					"pav:createdOn": "2025-05-09T00:21:34-07:00",
					"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"pav:lastUpdatedOn": "2025-05-09T00:39:36-07:00",
					"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"schema:schemaVersion": "1.6.0",
					"additionalProperties": false,
					"pav:version": "0.0.1",
					"bibo:status": "bibo:draft",
					"$schema": "http://json-schema.org/draft-04/schema#"
					}`,
			PhaseID: 1,
		},
		{
			Title:       "Data Collection / Acquisition",
			Description: "pending",
			Form: `{
					"@id": "https://repo.metadatacenter.org/templates/02b1f38a-6aa7-4e6d-b777-3cdc4d7df2d2",
					"@type": "https://schema.metadatacenter.org/core/Template",
					"@context": {
						"xsd": "http://www.w3.org/2001/XMLSchema#",
						"pav": "http://purl.org/pav/",
						"bibo": "http://purl.org/ontology/bibo/",
						"oslc": "http://open-services.net/ns/core#",
						"schema": "http://schema.org/",
						"schema:name": {
						"@type": "xsd:string"
						},
						"schema:description": {
						"@type": "xsd:string"
						},
						"pav:createdOn": {
						"@type": "xsd:dateTime"
						},
						"pav:createdBy": {
						"@type": "@id"
						},
						"pav:lastUpdatedOn": {
						"@type": "xsd:dateTime"
						},
						"oslc:modifiedBy": {
						"@type": "@id"
						}
					},
					"type": "object",
					"title": "Phase-2 template schema",
					"description": "Phase-2 template schema generated by the CEDAR Template Editor 2.7.7",
					"_ui": {
						"order": [
						"legitimacy",
						"consent",
						"privacy",
						"inequalities",
						"minimization"
						],
						"propertyLabels": {
						"legitimacy": "legitimacy",
						"consent": "consent",
						"privacy": "privacy",
						"inequalities": "inequalities",
						"minimization": "minimization"
						},
						"propertyDescriptions": {
						"legitimacy": "",
						"consent": "",
						"privacy": "",
						"inequalities": "",
						"minimization": ""
						}
					},
					"properties": {
						"@context": {
						"type": "object",
						"properties": {
							"rdfs": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2000/01/rdf-schema#"
							]
							},
							"xsd": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2001/XMLSchema#"
							]
							},
							"pav": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://purl.org/pav/"
							]
							},
							"schema": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://schema.org/"
							]
							},
							"oslc": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://open-services.net/ns/core#"
							]
							},
							"skos": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2004/02/skos/core#"
							]
							},
							"rdfs:label": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:isBasedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"schema:name": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:description": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"pav:derivedFrom": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:createdOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"pav:createdBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:lastUpdatedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"oslc:modifiedBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"skos:notation": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"legitimacy": {
							"enum": [
								"https://schema.metadatacenter.org/properties/4da0fa66-d34c-42f6-8db7-1318317c1427"
							]
							},
							"consent": {
							"enum": [
								"https://schema.metadatacenter.org/properties/c9bef392-5a1b-4adc-8b39-75d1afdb6afc"
							]
							},
							"privacy": {
							"enum": [
								"https://schema.metadatacenter.org/properties/bf317c27-7bfc-45a3-8041-7244a3ec8378"
							]
							},
							"inequalities": {
							"enum": [
								"https://schema.metadatacenter.org/properties/316e549a-28d9-4a06-81fe-39669c9e064f"
							]
							},
							"minimization": {
							"enum": [
								"https://schema.metadatacenter.org/properties/99b71216-f88b-4ead-aeb2-21e983209fd6"
							]
							}
						},
						"required": [
							"xsd",
							"pav",
							"schema",
							"oslc",
							"schema:isBasedOn",
							"schema:name",
							"schema:description",
							"pav:createdOn",
							"pav:createdBy",
							"pav:lastUpdatedOn",
							"oslc:modifiedBy"
						],
						"additionalProperties": false
						},
						"@id": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"@type": {
						"oneOf": [
							{
							"type": "string",
							"format": "uri"
							},
							{
							"type": "array",
							"minItems": 1,
							"items": {
								"type": "string",
								"format": "uri"
							},
							"uniqueItems": true
							}
						]
						},
						"schema:isBasedOn": {
						"type": "string",
						"format": "uri"
						},
						"schema:name": {
						"type": "string",
						"minLength": 1
						},
						"schema:description": {
						"type": "string"
						},
						"pav:derivedFrom": {
						"type": "string",
						"format": "uri"
						},
						"pav:createdOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"pav:createdBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"pav:lastUpdatedOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"oslc:modifiedBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"legitimacy": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "legitimacy field schema",
						"description": "legitimacy field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "legitimacy",
						"schema:description": "",
						"pav:createdOn": "2025-06-12T03:42:25-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-12T03:42:25-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Is the data source appropriate, legitimate, and transparent?",
						"@id": "https://repo.metadatacenter.org/template-fields/8f586100-5e5b-4a11-ba6a-64d2b89f2ccf",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"consent": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "consent field schema",
						"description": "consent field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "consent",
						"schema:description": "",
						"pav:createdOn": "2025-06-12T03:42:25-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-12T03:42:25-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Has informed consent been obtained where necessary?",
						"@id": "https://repo.metadatacenter.org/template-fields/8d83abe5-4d0f-43c2-a22f-74fb68e0d527",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"privacy": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "privacy field schema",
						"description": "privacy field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "privacy",
						"schema:description": "",
						"pav:createdOn": "2025-06-12T03:42:25-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-12T03:42:25-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Are there privacy risks or surveillance concerns?",
						"@id": "https://repo.metadatacenter.org/template-fields/a04ca880-3645-4383-93da-b6ace1e0fee1",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"inequalities": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "inequalities field schema",
						"description": "inequalities field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": true,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "inequalities",
						"schema:description": "",
						"pav:createdOn": "2025-06-12T03:42:25-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-12T03:42:25-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Could the data reflect or reinforce structural inequalities?",
						"@id": "https://repo.metadatacenter.org/template-fields/e23cff2e-b573-40d8-8b7d-a904beea417e",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"minimization": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "minimization field schema",
						"description": "minimization field schema generated by the CEDAR Template Editor 2.7.7",
						"_ui": {
							"inputType": "radio"
						},
						"_valueConstraints": {
							"requiredValue": false,
							"multipleChoice": false,
							"literals": [
							{
								"label": "Yes"
							},
							{
								"label": "No"
							}
							]
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "minimization",
						"schema:description": "",
						"pav:createdOn": "2025-06-12T03:42:25-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-06-12T03:42:25-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Are we collecting only the data we need (data minimization)?",
						"@id": "https://repo.metadatacenter.org/template-fields/29961309-6b9d-4148-ac08-48a804f7ca61",
						"$schema": "http://json-schema.org/draft-04/schema#"
						}
					},
					"required": [
						"@context",
						"@id",
						"schema:isBasedOn",
						"schema:name",
						"schema:description",
						"pav:createdOn",
						"pav:createdBy",
						"pav:lastUpdatedOn",
						"oslc:modifiedBy",
						"legitimacy",
						"consent",
						"privacy",
						"inequalities",
						"minimization"
					],
					"schema:name": "phase-2",
					"schema:description": "",
					"pav:createdOn": "2025-06-12T03:42:25-07:00",
					"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"pav:lastUpdatedOn": "2025-06-12T03:42:25-07:00",
					"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"schema:schemaVersion": "1.6.0",
					"additionalProperties": false,
					"pav:version": "0.0.1",
					"bibo:status": "bibo:draft",
					"$schema": "http://json-schema.org/draft-04/schema#"
					}`,
			PhaseID: 2,
		},
	}
	return db.Create(&reflections).Error
}

func (s ReflectionSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM reflections").Error
}

// ReflectionAnswer

type ReflectionAnswerSeeder struct{}

func (s ReflectionAnswerSeeder) Seed(db *gorm.DB) error {
	answers := []models.ReflectionAnswer{}
	return db.Create(&answers).Error
}

func (s ReflectionAnswerSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM reflection_answers").Error
}

// Journal

type JournalSeeder struct{}

func (s JournalSeeder) Seed(db *gorm.DB) error {
	journals := []models.Journal{
		{
			Title:       "",
			Description: "",
			Form: `{
					"@id": "https://repo.metadatacenter.org/templates/170e74d2-e93d-4a54-9637-27aca7af14ca",
					"@type": "https://schema.metadatacenter.org/core/Template",
					"@context": {
						"xsd": "http://www.w3.org/2001/XMLSchema#",
						"pav": "http://purl.org/pav/",
						"bibo": "http://purl.org/ontology/bibo/",
						"oslc": "http://open-services.net/ns/core#",
						"schema": "http://schema.org/",
						"schema:name": {
						"@type": "xsd:string"
						},
						"schema:description": {
						"@type": "xsd:string"
						},
						"pav:createdOn": {
						"@type": "xsd:dateTime"
						},
						"pav:createdBy": {
						"@type": "@id"
						},
						"pav:lastUpdatedOn": {
						"@type": "xsd:dateTime"
						},
						"oslc:modifiedBy": {
						"@type": "@id"
						}
					},
					"type": "object",
					"title": "Journal template schema",
					"description": "Journal template schema generated by the CEDAR Template Editor 2.7.1",
					"_ui": {
						"order": [
						"actions-decisions",
						"guidelines-tools-methods",
						"notes-questions"
						],
						"propertyLabels": {
						"actions-decisions": "actions-decisions",
						"guidelines-tools-methods": "guidelines-tools-methods",
						"notes-questions": "notes-questions"
						},
						"propertyDescriptions": {
						"actions-decisions": "",
						"guidelines-tools-methods": "",
						"notes-questions": ""
						}
					},
					"properties": {
						"@context": {
						"type": "object",
						"properties": {
							"rdfs": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2000/01/rdf-schema#"
							]
							},
							"xsd": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2001/XMLSchema#"
							]
							},
							"pav": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://purl.org/pav/"
							]
							},
							"schema": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://schema.org/"
							]
							},
							"oslc": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://open-services.net/ns/core#"
							]
							},
							"skos": {
							"type": "string",
							"format": "uri",
							"enum": [
								"http://www.w3.org/2004/02/skos/core#"
							]
							},
							"rdfs:label": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:isBasedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"schema:name": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"schema:description": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"pav:derivedFrom": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:createdOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"pav:createdBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"pav:lastUpdatedOn": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:dateTime"
								]
								}
							}
							},
							"oslc:modifiedBy": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"@id"
								]
								}
							}
							},
							"skos:notation": {
							"type": "object",
							"properties": {
								"@type": {
								"type": "string",
								"enum": [
									"xsd:string"
								]
								}
							}
							},
							"actions-decisions": {
							"enum": [
								"https://schema.metadatacenter.org/properties/e41c211e-c28b-48d0-9005-95ab7245e46c"
							]
							},
							"guidelines-tools-methods": {
							"enum": [
								"https://schema.metadatacenter.org/properties/a33f7259-20d8-4f42-aa90-ff2ffbbf4f8f"
							]
							},
							"notes-questions": {
							"enum": [
								"https://schema.metadatacenter.org/properties/ebc425cc-1ce1-476e-9149-3c3cb77aba15"
							]
							}
						},
						"required": [
							"xsd",
							"pav",
							"schema",
							"oslc",
							"schema:isBasedOn",
							"schema:name",
							"schema:description",
							"pav:createdOn",
							"pav:createdBy",
							"pav:lastUpdatedOn",
							"oslc:modifiedBy"
						],
						"additionalProperties": false
						},
						"@id": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"@type": {
						"oneOf": [
							{
							"type": "string",
							"format": "uri"
							},
							{
							"type": "array",
							"minItems": 1,
							"items": {
								"type": "string",
								"format": "uri"
							},
							"uniqueItems": true
							}
						]
						},
						"schema:isBasedOn": {
						"type": "string",
						"format": "uri"
						},
						"schema:name": {
						"type": "string",
						"minLength": 1
						},
						"schema:description": {
						"type": "string"
						},
						"pav:derivedFrom": {
						"type": "string",
						"format": "uri"
						},
						"pav:createdOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"pav:createdBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"pav:lastUpdatedOn": {
						"type": [
							"string",
							"null"
						],
						"format": "date-time"
						},
						"oslc:modifiedBy": {
						"type": [
							"string",
							"null"
						],
						"format": "uri"
						},
						"actions-decisions": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "actions-decisions field schema",
						"description": "actions-decisions field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "textarea"
						},
						"_valueConstraints": {
							"requiredValue": true
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "actions-decisions",
						"schema:description": "",
						"pav:createdOn": "2025-05-20T01:06:40-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-05-20T01:06:40-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Actions / Decisions",
						"@id": "https://repo.metadatacenter.org/template-fields/fe584866-54e8-4946-b72c-8148eaacd826",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"guidelines-tools-methods": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "guidelines-tools-methods field schema",
						"description": "guidelines-tools-methods field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "textarea"
						},
						"_valueConstraints": {
							"requiredValue": true
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "guidelines-tools-methods",
						"schema:description": "",
						"pav:createdOn": "2025-05-20T01:06:40-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-05-20T01:06:40-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Guidelines / Tools / Methods Used",
						"@id": "https://repo.metadatacenter.org/template-fields/07c7f9d6-6d47-4415-98f2-0510c1dc37f8",
						"$schema": "http://json-schema.org/draft-04/schema#"
						},
						"notes-questions": {
						"@type": "https://schema.metadatacenter.org/core/TemplateField",
						"@context": {
							"xsd": "http://www.w3.org/2001/XMLSchema#",
							"pav": "http://purl.org/pav/",
							"bibo": "http://purl.org/ontology/bibo/",
							"oslc": "http://open-services.net/ns/core#",
							"schema": "http://schema.org/",
							"skos": "http://www.w3.org/2004/02/skos/core#",
							"schema:name": {
							"@type": "xsd:string"
							},
							"schema:description": {
							"@type": "xsd:string"
							},
							"skos:prefLabel": {
							"@type": "xsd:string"
							},
							"skos:altLabel": {
							"@type": "xsd:string"
							},
							"pav:createdOn": {
							"@type": "xsd:dateTime"
							},
							"pav:createdBy": {
							"@type": "@id"
							},
							"pav:lastUpdatedOn": {
							"@type": "xsd:dateTime"
							},
							"oslc:modifiedBy": {
							"@type": "@id"
							}
						},
						"type": "object",
						"title": "notes-questions field schema",
						"description": "notes-questions field schema generated by the CEDAR Template Editor 2.7.1",
						"_ui": {
							"inputType": "textarea"
						},
						"_valueConstraints": {
							"requiredValue": true
						},
						"properties": {
							"@type": {
							"oneOf": [
								{
								"type": "string",
								"format": "uri"
								},
								{
								"type": "array",
								"minItems": 1,
								"items": {
									"type": "string",
									"format": "uri"
								},
								"uniqueItems": true
								}
							]
							},
							"@value": {
							"type": [
								"string",
								"null"
							]
							},
							"rdfs:label": {
							"type": [
								"string",
								"null"
							]
							}
						},
						"required": [
							"@value"
						],
						"schema:name": "notes-questions",
						"schema:description": "",
						"pav:createdOn": "2025-05-20T01:06:40-07:00",
						"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"pav:lastUpdatedOn": "2025-05-20T01:06:40-07:00",
						"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
						"schema:schemaVersion": "1.6.0",
						"additionalProperties": false,
						"skos:prefLabel": "Notes / Open Questions",
						"@id": "https://repo.metadatacenter.org/template-fields/270781a8-2fa4-4481-85e6-541dd85356ab",
						"$schema": "http://json-schema.org/draft-04/schema#"
						}
					},
					"required": [
						"@context",
						"@id",
						"schema:isBasedOn",
						"schema:name",
						"schema:description",
						"pav:createdOn",
						"pav:createdBy",
						"pav:lastUpdatedOn",
						"oslc:modifiedBy",
						"actions-decisions",
						"guidelines-tools-methods",
						"notes-questions"
					],
					"schema:name": "Journal",
					"schema:description": "",
					"pav:createdOn": "2025-05-20T01:05:00-07:00",
					"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"pav:lastUpdatedOn": "2025-05-20T01:06:40-07:00",
					"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"schema:schemaVersion": "1.6.0",
					"additionalProperties": false,
					"pav:version": "0.0.1",
					"bibo:status": "bibo:draft",
					"$schema": "http://json-schema.org/draft-04/schema#"
					}`,
			PhaseID: 1,
		},
		{
			Title:       "",
			Description: "",
			Form: `{
							"@id": "https://repo.metadatacenter.org/templates/198eb3c4-4b2a-43e4-b799-e681b944cccd",
							"@type": "https://schema.metadatacenter.org/core/Template",
							"@context": {
								"xsd": "http://www.w3.org/2001/XMLSchema#",
								"pav": "http://purl.org/pav/",
								"bibo": "http://purl.org/ontology/bibo/",
								"oslc": "http://open-services.net/ns/core#",
								"schema": "http://schema.org/",
								"schema:name": {
								"@type": "xsd:string"
								},
								"schema:description": {
								"@type": "xsd:string"
								},
								"pav:createdOn": {
								"@type": "xsd:dateTime"
								},
								"pav:createdBy": {
								"@type": "@id"
								},
								"pav:lastUpdatedOn": {
								"@type": "xsd:dateTime"
								},
								"oslc:modifiedBy": {
								"@type": "@id"
								}
							},
							"type": "object",
							"title": "Phase-2-journal template schema",
							"description": "Phase-2-journal template schema generated by the CEDAR Template Editor 2.7.7",
							"_ui": {
								"order": [
								"notes"
								],
								"propertyLabels": {
								"notes": "notes"
								},
								"propertyDescriptions": {
								"notes": ""
								}
							},
							"properties": {
								"@context": {
								"type": "object",
								"properties": {
									"rdfs": {
									"type": "string",
									"format": "uri",
									"enum": [
										"http://www.w3.org/2000/01/rdf-schema#"
									]
									},
									"xsd": {
									"type": "string",
									"format": "uri",
									"enum": [
										"http://www.w3.org/2001/XMLSchema#"
									]
									},
									"pav": {
									"type": "string",
									"format": "uri",
									"enum": [
										"http://purl.org/pav/"
									]
									},
									"schema": {
									"type": "string",
									"format": "uri",
									"enum": [
										"http://schema.org/"
									]
									},
									"oslc": {
									"type": "string",
									"format": "uri",
									"enum": [
										"http://open-services.net/ns/core#"
									]
									},
									"skos": {
									"type": "string",
									"format": "uri",
									"enum": [
										"http://www.w3.org/2004/02/skos/core#"
									]
									},
									"rdfs:label": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"xsd:string"
										]
										}
									}
									},
									"schema:isBasedOn": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"@id"
										]
										}
									}
									},
									"schema:name": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"xsd:string"
										]
										}
									}
									},
									"schema:description": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"xsd:string"
										]
										}
									}
									},
									"pav:derivedFrom": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"@id"
										]
										}
									}
									},
									"pav:createdOn": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"xsd:dateTime"
										]
										}
									}
									},
									"pav:createdBy": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"@id"
										]
										}
									}
									},
									"pav:lastUpdatedOn": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"xsd:dateTime"
										]
										}
									}
									},
									"oslc:modifiedBy": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"@id"
										]
										}
									}
									},
									"skos:notation": {
									"type": "object",
									"properties": {
										"@type": {
										"type": "string",
										"enum": [
											"xsd:string"
										]
										}
									}
									},
									"notes": {
									"enum": [
										"https://schema.metadatacenter.org/properties/66851c92-832c-4e29-8ac1-06029e90465a"
									]
									}
								},
								"required": [
									"xsd",
									"pav",
									"schema",
									"oslc",
									"schema:isBasedOn",
									"schema:name",
									"schema:description",
									"pav:createdOn",
									"pav:createdBy",
									"pav:lastUpdatedOn",
									"oslc:modifiedBy"
								],
								"additionalProperties": false
								},
								"@id": {
								"type": [
									"string",
									"null"
								],
								"format": "uri"
								},
								"@type": {
								"oneOf": [
									{
									"type": "string",
									"format": "uri"
									},
									{
									"type": "array",
									"minItems": 1,
									"items": {
										"type": "string",
										"format": "uri"
									},
									"uniqueItems": true
									}
								]
								},
								"schema:isBasedOn": {
								"type": "string",
								"format": "uri"
								},
								"schema:name": {
								"type": "string",
								"minLength": 1
								},
								"schema:description": {
								"type": "string"
								},
								"pav:derivedFrom": {
								"type": "string",
								"format": "uri"
								},
								"pav:createdOn": {
								"type": [
									"string",
									"null"
								],
								"format": "date-time"
								},
								"pav:createdBy": {
								"type": [
									"string",
									"null"
								],
								"format": "uri"
								},
								"pav:lastUpdatedOn": {
								"type": [
									"string",
									"null"
								],
								"format": "date-time"
								},
								"oslc:modifiedBy": {
								"type": [
									"string",
									"null"
								],
								"format": "uri"
								},
								"notes": {
								"@type": "https://schema.metadatacenter.org/core/TemplateField",
								"@context": {
									"xsd": "http://www.w3.org/2001/XMLSchema#",
									"pav": "http://purl.org/pav/",
									"bibo": "http://purl.org/ontology/bibo/",
									"oslc": "http://open-services.net/ns/core#",
									"schema": "http://schema.org/",
									"skos": "http://www.w3.org/2004/02/skos/core#",
									"schema:name": {
									"@type": "xsd:string"
									},
									"schema:description": {
									"@type": "xsd:string"
									},
									"skos:prefLabel": {
									"@type": "xsd:string"
									},
									"skos:altLabel": {
									"@type": "xsd:string"
									},
									"pav:createdOn": {
									"@type": "xsd:dateTime"
									},
									"pav:createdBy": {
									"@type": "@id"
									},
									"pav:lastUpdatedOn": {
									"@type": "xsd:dateTime"
									},
									"oslc:modifiedBy": {
									"@type": "@id"
									}
								},
								"type": "object",
								"title": "notes field schema",
								"description": "notes field schema generated by the CEDAR Template Editor 2.7.7",
								"_ui": {
									"inputType": "textarea"
								},
								"_valueConstraints": {
									"requiredValue": true
								},
								"properties": {
									"@type": {
									"oneOf": [
										{
										"type": "string",
										"format": "uri"
										},
										{
										"type": "array",
										"minItems": 1,
										"items": {
											"type": "string",
											"format": "uri"
										},
										"uniqueItems": true
										}
									]
									},
									"@value": {
									"type": [
										"string",
										"null"
									]
									},
									"rdfs:label": {
									"type": [
										"string",
										"null"
									]
									}
								},
								"required": [
									"@value"
								],
								"schema:name": "notes",
								"schema:description": "",
								"pav:createdOn": "2025-06-12T03:49:36-07:00",
								"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
								"pav:lastUpdatedOn": "2025-06-12T03:49:36-07:00",
								"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
								"schema:schemaVersion": "1.6.0",
								"additionalProperties": false,
								"skos:prefLabel": "Notes",
								"@id": "https://repo.metadatacenter.org/template-fields/4c4d5076-470c-4a19-8ecf-b740255b7b77",
								"$schema": "http://json-schema.org/draft-04/schema#"
								}
							},
							"required": [
								"@context",
								"@id",
								"schema:isBasedOn",
								"schema:name",
								"schema:description",
								"pav:createdOn",
								"pav:createdBy",
								"pav:lastUpdatedOn",
								"oslc:modifiedBy",
								"notes"
							],
							"schema:name": "phase-2-journal",
							"schema:description": "",
							"pav:createdOn": "2025-06-12T03:49:36-07:00",
							"pav:createdBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"pav:lastUpdatedOn": "2025-06-12T03:49:36-07:00",
					"oslc:modifiedBy": "https://metadatacenter.org/users/d268159d-2c15-41cd-8a63-0f822fb56d26",
					"schema:schemaVersion": "1.6.0",
					"additionalProperties": false,
					"pav:version": "0.0.1",
					"bibo:status": "bibo:draft",
					"$schema": "http://json-schema.org/draft-04/schema#"
					}`,
			PhaseID: 2,
		},
	}
	return db.Create(&journals).Error
}

func (s JournalSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM journal").Error
}

// JournalAnswer

type JournalAnswerSeeder struct{}

func (s JournalAnswerSeeder) Seed(db *gorm.DB) error {
	answers := []models.JournalAnswer{}
	return db.Create(&answers).Error
}

func (s JournalAnswerSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM journal_answers").Error
}

// Recommendation

type RecommendationSeeder struct{}

func (s RecommendationSeeder) Seed(db *gorm.DB) error {
	recommendations := []models.Recommendation{
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 0,
		},
		{
			ReflectionID:     1,
			ToolID:           2,
			BinaryEvaluation: 0,
		},
		{
			ReflectionID:     1,
			ToolID:           3,
			BinaryEvaluation: 0,
		},
		{
			ReflectionID:     1,
			ToolID:           4,
			BinaryEvaluation: 0,
		},
		{
			ReflectionID:     1,
			ToolID:           5,
			BinaryEvaluation: 0,
		},
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 1,
		},
		{
			ReflectionID:     1,
			ToolID:           2,
			BinaryEvaluation: 2,
		},
		{
			ReflectionID:     1,
			ToolID:           3,
			BinaryEvaluation: 4,
		},
		{
			ReflectionID:     1,
			ToolID:           4,
			BinaryEvaluation: 8,
		},
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 3,
		},
		{
			ReflectionID:     1,
			ToolID:           2,
			BinaryEvaluation: 3,
		},
		{
			ReflectionID:     1,
			ToolID:           3,
			BinaryEvaluation: 5,
		},
		{
			ReflectionID:     1,
			ToolID:           4,
			BinaryEvaluation: 5,
		},
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 6,
		},
		{
			ReflectionID:     1,
			ToolID:           3,
			BinaryEvaluation: 6,
		},
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 7,
		},
		{
			ReflectionID:     1,
			ToolID:           4,
			BinaryEvaluation: 9,
		},
		{
			ReflectionID:     1,
			ToolID:           3,
			BinaryEvaluation: 10,
		},
		{
			ReflectionID:     1,
			ToolID:           4,
			BinaryEvaluation: 10,
		},
		{
			ReflectionID:     1,
			ToolID:           5,
			BinaryEvaluation: 10,
		},
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 11,
		},
		{
			ReflectionID:     1,
			ToolID:           3,
			BinaryEvaluation: 11,
		},
		{
			ReflectionID:     1,
			ToolID:           4,
			BinaryEvaluation: 11,
		},
		{
			ReflectionID:     1,
			ToolID:           2,
			BinaryEvaluation: 12,
		},
		{
			ReflectionID:     1,
			ToolID:           1,
			BinaryEvaluation: 13,
		},
		{
			ReflectionID:     1,
			ToolID:           5,
			BinaryEvaluation: 13,
		},
		{
			ReflectionID:     1,
			ToolID:           5,
			BinaryEvaluation: 15,
		},
		{
			ReflectionID:     2,
			ToolID:           16,
			BinaryEvaluation: 0,
		},
	}
	return db.Create(&recommendations).Error
}

func (s RecommendationSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM recommendations").Error
}

// JournalAnswer

type RecommendationAnswerSeeder struct{}

func (s RecommendationAnswerSeeder) Seed(db *gorm.DB) error {
	answers := []models.JournalAnswer{}
	return db.Create(&answers).Error
}

func (s RecommendationAnswerSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM recommendation_answers").Error
}
