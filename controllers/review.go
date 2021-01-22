package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/theToulBox/ResuCan/views"
)

var (
	Soft = []string{"Clarity", "Confidence", "Respect", "Empathy", "Listening", "Verbal communication", "Non-verbal communication", "Written communication", "Constructive feedback", "Friendliness", "Conflict management", "Delegation", "Active listening", "Collaboration", "Cooperation", "Coordination", "Idea exchange", "Mediation", "Negotiating", "Self-management", "Decision-making", "Calmness", "Optimism", "Open-mindedness", "Analysis", "Self-confidence", "Organization", "Self-motivation", "Lateral thinking", "Logical reasoning", "Initiative", "Persistence", "Observation", "Persuasion", "Negotiation", "Brainstorming", "Decision making", "Divergent thinking", "Inspiration", "Imagination", "Reframing", "Mind mapping", "Insight", "Innovation", "Experimenting", "Questioning", "Design", "Integrity", "Responsibility", "Discipline", "Initiative", "Dependability", "Commitment", "Self-motivated", "Professionalism", "Teamwork", "Time-management", "Empathy", "Humor", "Mentoring", "Networking", "Sensitivity", "Patience", "Tolerance", "Public speaking", "Positive reinforcement", "Diplomacy", "Resourcefulness", "Thinking outside the box", "Tolerance of change", "uncertainty", "Troubleshooting", "Value education", "Willingness to learn", "Human Capital", "Dispute resolution", "Facilitation", "Giving clear feedback", "Inspiring people", "Managing difficult conversations", "Managing remote/virtual teams", "Meeting management", "Mentoring", "Successful coaching", "Supervising", "Talent management", "Confidence", "Cooperation", "Courtesy", "Energy", "Enthusiasts", "Friendliness", "Honesty", "Dealing with difficult situations", "Dealing with office politics", "Disability awareness", "Diversity awareness", "Empathy", "Establishing interpersonal relationships", "Dealing with difficult personalities", "Intercultural competence", "Influence", "Networking", "Persuasion", "Self-awareness", "Attentiveness", "Business ethics", "Competitiveness", "Dedication", "Dependability", "Following direction", "Independence", "Meeting deadlines", "Motivation", "Perseverance", "Persistence", "Planning", "Proper business etiquette", "Punctuality", "Reliability", "Resilience", "Results-orientation", "Scheduling", "Self-direction", "Self-monitoring", "Self-supervising", "Staying on task", "Prioritizing", "Self-starter", "Planning", "Decision making", "Focus", "Delegation", "Stress management", "Coping", "Organization", "Project management", "Selflessness", "Agility", "Humility", "Cultural intelligence", "Authenticity", "Versatility", "Generosity", "Trust", "Introspection", "Memory", "Acuity", "Recall", "Questioning"}
	Hard = []string{"Software Development", "Automation", "Test Automation", "Airtable", "Mongo", "Dynamodb", "Dynamo", "Postgres", "Microsoft Office", "Social media", "HTML", "Analytics", "Pivot tables", "Foreign language", "Digital communication", "Copywriting", "Data mining", "Data presentation", "Resource management", " Data engineering", "Database management", "Search Engine Optimization", "SEO", "Search Engine Marketing", "SEM", "Marketing Campaign Management", "Google Analytics", "CMS", "WordPress", "Scrum", "Agile", "Trello", "Zoho", "User Interface", "UI/UX", "User interaction", "user experience", "Design", "Adobe Creative Suite", "Photoshop", "InDesign", "Big Data", "Programming", "Coding", "Java", "C++", "Python", "GO", "GoLang", "SQL", "Database", "InVision", "Zeppelin", "Machinery", "Operations", "Automation", "Robotics Process Automation", "Adobe", "SME", "Segment Reporting", "Reporting", "Customer relationship management", "SAP", "Accounting", "Finance", "Booking", "Sales", "Analytics", "HIPA", "Web Analytics", "HTML", "Wordpress", "Email marketing", "Web scraping", "chief risk officer", "A/B Testing", "Data visualization", "Keyword Optimization", "Project Management", "Campaign Management", "Social Media Marketing", "Mobile Marketing", "B2B Marketing", "Business to Business Marketing", "Consumer Behavior", "Brand Management", "Copywriting", "Storytelling", "Data Analytics", "Trend Analysis", "CMS Tools", "Task Delegation", "Business Development", "Strategic Management", "Planning", "Negotiation", "Proposal Writing", "Problem-solving", "Innovation", "Six Sigma", "Customer Relationship Management", "Cold-calling", "Public Speaking", "Closing", "Lead Generation", "Buyer-Responsive Selling", "Buyer Engagement", "Product Knowledge", "Persuasion", "Ad Design", "Illustrator", "InDesign", "Architecture", "Storage", "Data Management", "Networking Communication", "Middleware", "Application", "JSON", "RPC", "Oracle", "Software Revision Control Systems", "Android", "iOS", "Web", "Angular", "Node", "Javascript", "Prontend", "CSS", "Administration", "Drafting", "Operating systems", "Quality assurance", "QA", "AWS", "AZURE", "GCP", "Data analytics", "Network", "ERP", "Enterprise resource planning", "Email marketing", "Web scraping", "Testing", "Kubernetes", "Helm", "A/B Testing", "Data visualization", "Pattern-finding", "Keyword Optimization", "Social media", "Mobile marketing", "Social media ads", "B2B Marketing", "Brand management", "Copywriting", "Word", "Excel", "Access", "Publisher", "Outlook", "Powerpoint", "Sharepoint", "Filing", "Data entry", "Bookkeeping", "TurboTax", "Intuit", "Research", "Data analysis", "Technical writing", "Big Data Analysis", "Cognos Analytics", "Visual Basic", "Hive", "Python", "Scala", "Matlab", "NET", "STATA", "SPSS", "SAS", "Data Mapping", "Entity Relationship Diagrams", "Wireframes", "Big Data", "Machine learning", "System Context Diagrams", "CPR", "Patient care", "Advanced Cardiac Life Support", "ACLS", "Telemetry", "C#", "Groovy", "Selenium", "Budgeting", "Business Planning", "Business Re-engineering", "Change Management", "Consolidation", "Cost Control", "Decision Making", "Developing Policies", "Diversification", "Employee Evaluation", "Financing", "Government Relations", "Hiring", "International Management", "Investor Relations", "IPO", "Joint Ventures", "Labour Relations", "Merger & Acquisitions", "Multi-sites Management", "Negotiation", "Profit & Loss", "Organizational Development", "Project Management", "Staff Development", "Strategic Partnership", "Strategic Planning Supervision", "Bidding", "Call Centre Operations", "Continuous Improvement", "Contract Management", "Environmental Protection", "Facility Management", "Inventory Control", "Manpower Planning", "Operations Research", "Outsourcing", "Policies & Procedures", "Project Co-ordination", "Project Management", "Equipment Design", "Equipment Maintenance & Repair", "Equipment Management", "ISO", "TQM", "Order Processing", "Plant Design & Layout", "Process Engineering", "Production Planning", "Safety Engineering", "Transportation", "JIT Purchasing & Procurement Shipping", "Traffic Management Warehousing", "Research & Development", "Design & Specification", "Diagnostics", "Feasibility Studies", "Field Studies", "Lab Management", "Lab Design", "New Equipment Design", "Patent Application", "Product Development", "Product Testing", "Prototype Development", "Statistical Analysis", "Contract Negotiation", "Customer Relations", "Customer Service", "Direct Sales", "Distributor Relations", "E-Commerce", "Forecasting Incentive Programs", "International Business Development", "International Expansion", "New Account Development", "Proposal Writing", "Product Demonstrations", "Telemarketing", "Trade Shows", "Sales Administration", "Sales Analysis", "Sales Kits", "Sales Management", "Salespersons Recruitment", "Show Room Management", "Sales Support", "Sales Training", "Advertising", "Brand Management", "Channel Marketing", "Competitive Analysis", "Copywriting", "Corporate Identity", "Image Development", "Logo Development", "Market Research", "Market Analysis", "Marketing Communication", "Marketing Plan", "Marketing promotions", "Media Buying", "Media Relations", "Merchandising", "New Product Development", "Online Marketing", "Packaging Pricing", "Product Launch Administration", "Contract Negotiation", "Equipment Purchasing", "Forms and Methods Leases", "Mailroom Management", "Office Management", "Policies & Procedures", "Reception", "Records Management", "Security", "Space Planning", "Word Processing", "Contract Preparation", "Copyrights & Trademarks", "Corporate Law", "Company Secretary", "Employment Ordinance", "Intellectual Property", "International Agreements", "Licensing", "Mergers", "Patents", "Shareholder proxies", "Stock Administration", "Accounting Management", "Accounts Payable", "Venture Capital Relations", "Accounts Receivable", "Acquisitions & Mergers", "Actuarial", "Rating", "Auditing", "Banking Relation", "Budget Control", "Capital Budgeting", "Capital Investment", "Cash Management", "Cost Accounting", "Cost Control", "Credit Collections", "Debt Negotiations", "Debt Management", "Feasibility Studies", "Financial Analysis", "Financial Reporting", "Financing Forecasting", "Foreign Exchange", "General Ledger Insurance", "Internal Controls", "Investor Relations", "Lending", "Fund Management", "Profit Planning", "Risk Management", "Stockholder Relations", "Tax", "Treasury", "Investor Presentations", "Arbitration", "Mediation", "Career Counseling", "Career Coaching", "Classified Advertisements", "Company Orientation", "Workforce Forecast", "Compensation & Benefits", "Corporate Culture", "Training Administration", "Employee Discipline", "Employee Selection", "Executive Recruiting", "Grievance Resolution", "Human Resources Management", "Industrial Relations", "Job Analysis", "Labour Negotiations", "Outplacement", "Performance Appraisal", "Salary Administration", "Succession Planning", "Team Building", "Training", "Algorithm Development", "Application Database Administration", "Applications Development", "Business Systems Planning", "Web Site Editor", "Capacity Planning", "Machinery", "CAD", "EDI", "Enterprise Asset Management", "EAP", "Integration Software", "Software Customization", "System Analysis System Design", "System Development", "Technical Evangelism", "Technical Support", "Telecommunications", "Tracking System", "UNIX", "Usability Engineering", "User Education", "User Documentation", "User Interface Vendor Sourcing", "Voice", "Web Development", "Web design", "Web Site Content Writer", "Word Processing", "Development", "Dreamweaver", "Flash Freehand", "Figma", "Sketch", "Illustrator", "Photoshop", "Picasa", "Typography", "Public Relation", "B2B Communication", "Community Relations", "Speech Writing", "Corporate Image", "Corporate Philanthropy", "Corporate Publications", "Corporate Relations", "Employee Communication", "Event Planning", "Fund Raising", "Government Relations", "Investor Collateral", "Media Presentations", "Press Release"}
)

// Review contains the two views to be rendered
type Review struct {
	NewView    *views.View
	ResultView *views.View
}

// ReviewForm is the form to be completed to start an analysis
type ReviewForm struct {
	Resume      string `schema:"resume"`
	Description string `schema:"description"`
}

// Result is the conclusion of the analysis
type Result struct {
	LinkedIn     bool
	HardSkills   []string
	SoftSkills   []string
	ResumeLength int
}

// NewReview Renders the views.
func NewReview() *Review {
	return &Review{
		NewView:    views.NewView("css", "review/new"),
		ResultView: views.NewView("css", "review/result"),
	}
}

// New renders a review view.
func (re *Review) New(w http.ResponseWriter, r *http.Request) {
	re.NewView.Render(w, r, nil)
}

// Review will analyze the Resume against the Job Description
func (re *Review) Review(w http.ResponseWriter, r *http.Request) {
	// now i have the form data container from parsing the data
	vd := views.Data{}
	form := ReviewForm{}
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		re.NewView.Render(w, r, vd)
	}

	// Run the analysis
	res, err := re.Analyze(form.Resume, form.Description)
	if err != nil {
		vd.SetAlert(err)
		re.NewView.Render(w, r, vd)
	}

	// collect the res
	vd.Yield = res

	// render the view
	re.ResultView.Render(w, r, vd)

}

// Analyze contains the algorithm for performing the critique
func (re *Review) Analyze(r, d string) (*Result, error) {
	resu := strings.ToLower(r)
	desc := strings.ToLower(d)
	linkedIn := HasLinkedIn(resu)
	rLength := ResumeLength(resu)
	// Collect keywords from Job description
	JobDescrHardSkills := FindSkills(desc, Hard)
	fmt.Println("JobDescrHardSkills ", JobDescrHardSkills)

	JobDescrSoftSkills := FindSkills(desc, Soft)
	fmt.Println("JobDescrSoftSkills", JobDescrSoftSkills)

	// Collect keywords from resume
	ResumeHardSkills := FindSkills(resu, Hard)
	fmt.Println("ResumeHardSkills", ResumeHardSkills)

	ResumeSoftSkills := FindSkills(resu, Soft)
	fmt.Println("ResumeSoftSkills", ResumeSoftSkills)
	// If Resume does not have hard skills that were found in the job description
	// then add the Result as recommendation
	hSkills := Diff(JobDescrHardSkills, ResumeHardSkills)
	fmt.Println("Resume could use these hard skills", hSkills)
	// If Resume does not have soft skills that were found in the job description
	// then add the result as recommendation
	sSkills := Diff(JobDescrSoftSkills, ResumeSoftSkills)
	fmt.Println("Resume Could use these soft skills", sSkills)

	// mSkillCount := measurableSkillCount(t)

	res := &Result{
		LinkedIn:     linkedIn,
		HardSkills:   hSkills,
		SoftSkills:   sSkills,
		ResumeLength: rLength,
	}
	return res, nil
}

func FindSkills(t string, set []string) (skills []string) {
	for _, s := range set {
		s = strings.ToLower(s)
		if strings.Contains(t, s) {
			skills = append(skills, s)
		}
	}
	// handle the case when there aren't any matches
	if skills == nil {
		skills = append(skills, "")
	}
	return
}

func HasLinkedIn(t string) bool {
	return strings.Contains(t, "linkedin")
}

func Diff(a, b []string) []string {
	a = SortIfNeeded(a)
	b = SortIfNeeded(b)
	var d []string
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		c := strings.Compare(a[i], b[j])
		if c == 0 {
			i++
			j++
		} else if c < 0 {
			d = append(d, a[i])
			i++
		} else {
			d = append(d, b[j])
			j++
		}
	}
	d = append(d, a[i:len(a)]...)
	d = append(d, b[j:len(b)]...)
	d = RemoveDups(d)
	return d
}

func SortIfNeeded(a []string) []string {
	if sort.StringsAreSorted(a) {
		return a
	}
	s := append(a[:0:0], a...)
	sort.Strings(s)
	return s
}

func RemoveDups(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

func ResumeLength(res string) int {
	return len(strings.Split(res, " "))
}

// func MeasurableSkillCount(t string) [][]byte {
// 	re := regexp.MustCompile(`[0-9]%`)
// 	percentages := re.FindAll([]byte(t), -1)
// 	return percentages
// }
