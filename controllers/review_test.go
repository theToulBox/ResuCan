package controllers

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAnalyze(t *testing.T) {
	cases := []struct {
		desc    string
		resu    string
		jobPost string
		result  *Result
	}{
		{
			desc:    "TestAnalyzeBothSkillsMissing",
			resu:    "A 20USD resume is more complex than a $1.00 but not as much as a 20% fee",
			jobPost: "automation clarity",
			result: &Result{
				Score:        12.57,
				LinkedIn:     false,
				HardSkills:   []string{"automation"},
				SoftSkills:   []string{"clarity"},
				ResumeLength: 17,
				Measurable:   3,
			},
		},
		{
			desc:    "TestAnalyzeNoMatchInJobDescr",
			resu:    "resume is more complex",
			jobPost: "a poorly written job description",
			result: &Result{
				Score:        0.13,
				LinkedIn:     false,
				HardSkills:   nil,
				SoftSkills:   nil,
				ResumeLength: 4,
				Measurable:   0,
			},
		},
		{
			desc:    "TestAnalyzeExcellentMatch",
			resu:    "Kanye West www.linkedin.com\n---------\n\nEDUCATION\n---------\nB.S. Computer Science, University of Houston - Main Campus, December 2018\n\nSKILLS\n------\nPython, GO, Docker, Kubernetes, KubeBench, DockerBench, KubeHunter, KOPS, Helm, Terraform, Terragrunt, Anchore, Veracode, Machine Learning\n\nPERSONAL PROJECTS\n-----------------\nResuCan.com\n  * Open Source Resume Scanning Project built in GoLang and Deployed on AWS\n  * Built a DevSecOps Pipeline through GitHub Actions to reduce likelihood of exploits making it to production\n  * Implemented a WAF to further protect web application\n\nGO-DevSecOps-Pipeline\n  * Open Source repository that shows how to create a free DevSecOps Pipeline through GitHub Actions\n  * Added OWASP Dynamic Application Scanning\n  * Demonstrates a Dockerized GoLang application deployment to AWS ElasticBeanstalk\nTheToulBox.com\n  * Where I blog about technlogy\n\nEXPERIENCE\n----------\nDevOps Engineer, HP, Inc., Houston, TX, 2018-Present\n\nTerragrunt Infrastructure Pipeline\n  * Enabled 200+ developers to provision any piece of AWS infrastructure they needed 24/7 365 days a year\n  * Reduced infrastructure resource provisioning time from days to minutes\n\nAWS WAF Implementation\n  * Increased AWS Web Application Firewall (WAF) on AWS Application Load Balancers (ALB) from 0% to 100% in production\n  * Innersourced the WAF into a Terragrunt template that is applied across hundreds of AWS ALB's\n\nCompliance Reporting Platform\n  * Reduced report creation from one hour to automatically happening in seconds via Python Lambdas\n  * Helped secure SOC 2 and SOC 1 compliance as a result of 100% monthly reporting\n\nGolden Amazon Machine Image\n  * Used Docker Bench to create a golden Amazon Machine Image (AMI)\n  * AMI is now the basis of 200+ Kubernetes Nodes\n  * AMI has Security Agents Built in to ensure compliance of SOC 2 and 1\n\nKubernetes Zero Downtime Upgrades\n  * Used Kubebench to determine which vulnerabilities existed in current version of Kubernetes cluster\n  * Upgraded Kubernetes Cluster version from 1.11 to 1.13 with zero downtime for CVE vulnerability patches\n\nGater - Gated Check In tool\n  * GoLang CLI that enabled smart code coverage check in on 300+ repositories\n  * Increased Code coverage from 10-60% on average\n\nADM - Automation of Repository Creation & Enforcement\n  * Created a GoLang CLI to enforce branch protection rules on 300+ repositories\n  * Branch protections increased by 51%\n\nHONORS\n------\n  * Terry Foundation Scholar, 2015-2018\n  * Texas Space Grant, 2015\n\nLEADERSHIP\n----------\n * Led Alpha Lambda Xi to place 2nd out of 1300 internationally for research project as Executive Vice President\n * Led Fire Drone a 3-person start-up to semi finals at SHPE 2017 National Conference Hackathon analysis commitment design organization self-motivated troubleshooting access administration architecture azure design development erp excel hiring network operating systems operations outsourcing problem-solving rating sas scala software development storage",
			jobPost: "Skip to main content\nIndeed Home\nSign in\nDevOps Engineer\nMedvantx\nRemote\nRemote\nMedvantx\n\n17 reviews\nRead what people are saying about working here.\nJob details\n\nJob Type\nFull-time\nQualifications\n\nBachelor's (Preferred)\nLinux administration/troubleshooting.: 3 years (Preferred)\nAzure: 3 years (Preferred)\nFull Job Description\n\nGENERAL PURPOSE:\n\nThe DevOps Engineer’s role is to strategically design, implement, and maintain information systems architecture and cloud infrastructure that support core organizational functions and assure their high availability. This individual will gain organizational commitment for all systems and infrastructure plans, as well as evaluate and select all technologies required to complete those plans.\n\nRESPONSIBILITIES:\n\nYou will design and implement long-term strategic goals and short-term tactical plans for managing and maintaining enterprise systems including servers, storage, network communications, and data center facilities. You will work closely with DevOps and cloud infrastructure architects and engineers to design, implement and manage secure, scalable and reliable cloud infrastructure environments. You will ensure that proposed and existing systems architectures are aligned with organizational goals and objectives. Duties include:\n\n· Implement infrastructure best practices in areas such as CI, CD, performance, scalability, security, and availability.\n\n· Provide architectural expertise, direction, and assistance to Systems Administrators, Systems Analysts, Systems Engineers, and software development teams.\n\n· Manage architecture and development of internal enterprise software and customization projects. Work with assigned development group resources to complete such projects.\n\n· Oversee IT cloud, server and network infrastructure procurement, installation, maintenance, and backup/restore processes.\n\n· Develop, document, and communicate plans for investing in systems infrastructure, including analysis of cost reduction opportunities.\n\n· Document the company’s existing systems infrastructure and technology portfolio and make recommendations for improvements and/or alternatives.\n\n· Review new and existing systems design projects and procurement or outsourcing plans for compliance with standards and architectural plans.\n\n· Confer with end-users, clients, or senior management to define business requirements for complex systems and infrastructure development.\n\n· Develop, manage, and test a disaster recovery plan for SaaS outages, data center operations, LAN/WAN connectivity, and server-based desktop resources.\n\n· Responsible for implementation and maintenance of enterprise software applications and systems.\n\n· Responsible for day-to-day operations and management of the infrastructure of the data center environment in order to provide 24x7 critical infrastructure coverage and ensure continual data center operations.\n\n· Develop and manage a systems capacity plan.\n\nEXPERIENCE REQUIRED:\n\n· Excellent knowledge of service and application delivery, as well as successful service level agreement accomplishments.\n\n· 3-5 years of experience with Linux administration and troubleshooting.\n\n· 3-5 years of experience with container-based infrastructure in Azure (Docker, Kubernetes, AKS, etc)\n\n· Experience with server operating systems, infrastructure services, and client/server applications.\n\n· Strong experience with virtualization technologies and server-based computing platforms including Citrix Presentation Server (XenApp), VMware vSphere, and VMware View.\n\n· Experience with DevOps practices and Azure DevOps tools.\n\n· Experience with ticketing tools (Zendesk preferred).\n\n· Knowledge of applicable data privacy practices and laws.\n\n· Highly self-motivated and directed with precise attention to detail.\n\n· Good understanding of the organization’s goals and objectives.\n\n· Proven analytical and problem-solving abilities.\n\n· Ability to effectively prioritize and execute tasks in a high-pressure environment.\n\n· Ability to effectively communicate relevant IT infrastructure related information to superiors and colleagues.\n\n· Anticipate business needs and proposes alternative business solutions.\n\nJob Type: Full-time\n\nPay: $0.00 per year\n\nBenefits:\n\n401(k)\nDental insurance\nHealth insurance\nLife insurance\nPaid time off\nVision insurance\nSchedule:\n\nMonday to Friday\nEducation:\n\nBachelor's (Preferred)\nExperience:\n\nLinux administration/troubleshooting.: 3 years (Preferred)\nAzure: 3 years (Preferred)\nWork Location:\n\nFully Remote\nJust posted\nIf you require alternative methods of application or screening, you must approach the employer directly to request this as Indeed is not responsible for the employer's application process.\n\nReport job\n\nApply Now\nSave this job\nShare this job\nDevelopment Operations Engineer jobs in Remote\nJobs at Medvantx in Remote\nDevelopment Operations Engineer salaries in Remote\nHiring Lab\nCareer Advice\nBrowse Jobs\nBrowse Companies\nSalaries\nFind Certifications\nIndeed Events\nWork at Indeed\nCountries\nAbout\nHelp Center\n© 2021 Indeed\nDo Not Sell My Personal Information\nAccessibility at Indeed\nPrivacy Center\nCookies\nPrivacy\nTerms\n",
			result: &Result{
				Score:        88.34,
				LinkedIn:     true,
				HardSkills:   nil,
				SoftSkills:   nil,
				ResumeLength: 432,
				Measurable:   25,
			},
		},
	}
	for _, tc := range cases {
		n := Review{}
		got, _ := n.Analyze(tc.resu, tc.jobPost)
		if !cmp.Equal(got, tc.result) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.result)
		}
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		desc string
		resu string
		set  []string
		want []string
	}{
		{"TestFindsHardSkillInResume", "automation is fun", Hard, []string{"automation"}},
		{"TestFindSoftSkillInResume", "clarity is what we need", Soft, []string{"clarity"}},
		{"TestFindSkillsEmptyString", "", Soft, nil},
		{"TestFindSkillsWonkyCase", "ClArItY is what we need", Soft, []string{"clarity"}},
		{"TestFindEducationRequirment", "We require a Bachelor's Receipt or Master's Heck even a Phd", Education, []string{"bachelor's", "master's", "phd"}},
	}
	for _, tc := range cases {
		got := Find(tc.resu, tc.set)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}

func TestHasWord(t *testing.T) {
	cases := []struct {
		desc string
		txt  string
		want string
		has  bool
	}{
		{"TestHasLinkedIn", "a resume should have a linkedin", "linkedin", true},
		{"TestNoLinkedIn", "some one's resume", "master's", false},
	}
	for _, tc := range cases {
		got := HasWord(tc.txt, tc.want)
		if got != tc.has {
			t.Errorf("%s: got %v want %v", tc.desc, tc.txt, tc.want)
		}
	}
}

func TestDiff(t *testing.T) {
	cases := []struct {
		desc   string
		post   []string
		skills []string
		want   []string
	}{
		{"TestDiffHasOne", []string{"eating"}, []string{"C++"}, []string{"C++"}},
		{"TestDiffNone", []string{"automation"}, []string{"automation"}, nil},
		{"TestDiffEmptyString", []string{""}, []string{"Resume"}, []string{"Resume"}},
	}
	for _, tc := range cases {
		got := Diff(tc.skills, tc.post)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		desc  string
		words []string
		want  []string
	}{
		{"TestDupsRemoves", []string{"soap", "soap", "soap"}, []string{"soap"}},
	}
	for _, tc := range cases {
		got := RemoveDups(tc.words)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}

func TestMeasurableSkillCount(t *testing.T) {
	cases := []struct {
		desc string
		resu string
		want float64
	}{
		{"TestFindNums", "Increased Fight Club recruits by 20% $40 $70.12 ", 3},
	}
	for _, tc := range cases {
		got := MeasurableSkillCount(tc.resu)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}
