package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Articles struct {
	Id      int       `form:"-"`
	Title   string    `form:"title" required`
	Content string    `orm:";type(text)" form:"content,textarea"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type Profile struct {
	Id               int       `orm:"column(id);auto"`
	UserId           int       `orm:"column(user_id);null"`
	Firstname        string    `orm:"column(firstname);size(45);null"`
	Middlename       string    `orm:"column(middlename);size(45);null"`
	Surname          string    `orm:"column(surname);size(45);null"`
	Sex              string    `orm:"column(sex);size(45);null"`
	Birthdate        time.Time `orm:"column(birthdate);type(date);null"`
	Height           int       `orm:"column(height);null"`
	Weight           int       `orm:"column(weight);null"`
	Smoker           int       `orm:"column(smoker);null"`
	Avdrinks         int       `orm:"column(avdrinks);null"`
	Address1         string    `orm:"column(address1);size(90);null"`
	Address2         string    `orm:"column(address2);size(90);null"`
	Address3         string    `orm:"column(address3);size(90);null"`
	Suburb           string    `orm:"column(suburb);size(45);null"`
	State            string    `orm:"column(state);size(10);null"`
	Postcode         int       `orm:"column(postcode);null"`
	Email1           string    `orm:"column(email1);size(90);null"`
	Email2           string    `orm:"column(email2);size(90);null"`
	PhoneHome        string    `orm:"column(phone_home);size(45);null"`
	PhoneMob         string    `orm:"column(phone_mob);size(45);null"`
	Msn              string    `orm:"column(msn);size(90);null"`
	Twitter          string    `orm:"column(twitter);size(90);null"`
	Facebook         string    `orm:"column(facebook);size(90);null"`
	Wechat           string    `orm:"column(wechat);size(90);null"`
	Profession       string    `orm:"column(profession);size(90);null"`
	Profilecol       string    `orm:"column(profilecol);size(45);null"`
	Title            string    `orm:"column(title);size(45);null"`
	Child1Name       string    `orm:"column(child1_name);size(45);null"`
	Child1Birthdate  time.Time `orm:"column(child1_birthdate);type(date);null"`
	Child1Sex        string    `orm:"column(child1_sex);size(10);null"`
	Child2Name       string    `orm:"column(child2_name);size(45);null"`
	Child2Birthdate  time.Time `orm:"column(child2_birthdate);type(date);null"`
	Child2Sex        string    `orm:"column(child2_sex);size(10);null"`
	Child3Name       string    `orm:"column(child3_name);size(45);null"`
	Child3Birthdate  time.Time `orm:"column(child3_birthdate);type(date);null"`
	Child3Sex        string    `orm:"column(child3_sex);size(10);null"`
	Child4Name       string    `orm:"column(child4_name);size(45);null"`
	Child4Birthdatae time.Time `orm:"column(child4_birthdatae);type(date);null"`
	Child4Sex        string    `orm:"column(child4_sex);size(10);null"`
	Child5Name       string    `orm:"column(child5_name);size(45);null"`
	Child5Birthdate  time.Time `orm:"column(child5_birthdate);type(date);null"`
	Child5Sex        string    `orm:"column(child5_sex);size(10);null"`
	NoChildren       int       `orm:"column(no_children);null"`
	NoPets           int       `orm:"column(no_pets);null"`
	Pet1Name         string    `orm:"column(pet1_name);size(45);null"`
	Pet1Type         string    `orm:"column(pet1_type);size(45);null"`
	Pet1Sex          string    `orm:"column(pet1_sex);size(10);null"`
	Pet2Name         string    `orm:"column(pet2_name);size(45);null"`
	Pet2Type         string    `orm:"column(pet2_type);size(45);null"`
	Pet2Sex          string    `orm:"column(pet2_sex);size(10);null"`
	Pet3Name         string    `orm:"column(pet3_name);size(45);null"`
	Pet3Sex          string    `orm:"column(pet3_sex);size(10);null"`
	Pet3Type         string    `orm:"column(pet3_type);size(45);null"`
	LastUpdated      time.Time `orm:"column(last_updated);type(datetime);null"`
	OptionsRadios    string    `orm:"column(options_radios);size(191);null"`
	Position         string    `orm:"column(position);size(45);null"`
}

// Addprofile insert a new profile into database and returns
// last inserted Id on success.
func AddProfile(m *Profile) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetprofileById retrieves profile by Id. Returns error if
// Id doesn't exist
func GetProfileById(id int) (v *Profile, err error) {
	o := orm.NewOrm()
	v = &Profile{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllprofile retrieves all profile matches certain condition. Returns empty list if
// no records exist
func GetAllProfile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Profile))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Profile
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// Updateprofile updates profile by Id and returns error if
// the record to be updated doesn't exist
func UpdateProfileById(m *Profile) (err error) {
	o := orm.NewOrm()
	v := Profile{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// Deleteprofile deletes profile by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProfile(id int) (err error) {
	o := orm.NewOrm()
	v := Profile{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Profile{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (a *Articles) TableName() string {
	return "articles"
}

func (a *Profile) TableName() string {
	return "profiles"
}

type LifeInsurance struct {
	Id                                     int       `orm:"column(id);auto"`
	BatchId                                int       `orm:"column(batch_id);null"`
	Gender                                 string    `orm:"column(gender);size(2);null"`
	Smoker                                 int       `orm:"column(smoker);null"`
	LastSmoked                             int       `orm:"column(last_smoked);null"`
	Age                                    int       `orm:"column(age);null"`
	CoverAmount                            int       `orm:"column(cover_amount);null"`
	CountryResidence                       string    `orm:"column(country_residence);size(45);null"`
	ResidencyStatus                        string    `orm:"column(residency_status);size(45);null"`
	Height                                 int       `orm:"column(height);null"`
	Weight                                 int       `orm:"column(weight);null"`
	Cancer                                 int       `orm:"column(cancer);null"`
	Diabetes                               int       `orm:"column(diabetes);null"`
	BloodDisorder                          int       `orm:"column(blood_disorder);null"`
	CancerType                             string    `orm:"column(cancer_type);size(255);null"`
	CancerSpread                           int       `orm:"column(cancer_spread);null"`
	CancerClear10                          int       `orm:"column(cancer_clear_10);null"`
	DiabetesComplications                  int       `orm:"column(diabetes_complications);null"`
	BloodDisorderType                      string    `orm:"column(blood_disorder_type);size(255);null"`
	BloodPressure                          int       `orm:"column(blood_pressure);null"`
	Cholesterol                            int       `orm:"column(cholesterol);null"`
	HeartProblem                           int       `orm:"column(heart_problem);null"`
	BloodPressureTreatment                 string    `orm:"column(blood_pressure_treatment);size(255);null"`
	KidneyProblem                          int       `orm:"column(kidney_problem);null"`
	CholesterolTreatment                   string    `orm:"column(cholesterol_treatment);size(255);null"`
	CholesterolCheck12                     int       `orm:"column(cholesterol_check_12);null"`
	HeartProblemDetails                    string    `orm:"column(heart_problem_details);size(255);null"`
	GastroIntestinal                       int       `orm:"column(gastro_intestinal);null"`
	BladderProblem                         int       `orm:"column(bladder_problem);null"`
	LungProblem                            int       `orm:"column(lung_problem);null"`
	GastroProblemDetails                   string    `orm:"column(gastro_problem_details);size(255);null"`
	HerniaFurtherProblems                  int       `orm:"column(hernia_further_problems);null"`
	GastritisEpisodes                      int       `orm:"column(gastritis_episodes);null"`
	GallstonesRemoved                      int       `orm:"column(gallstones_removed);null"`
	GallstonesOngoing                      int       `orm:"column(gallstones_ongoing);null"`
	UlcerHospitialised                     int       `orm:"column(ulcer_hospitialised);null"`
	GastroOtherTreatment                   int       `orm:"column(gastro_other_treatment);null"`
	GastroProblems5y                       int       `orm:"column(gastro_problems_5y);null"`
	KydneyBladderProblemDetails            string    `orm:"column(kydney_bladder_problem_details);size(255);null"`
	LungProblemDetails                     string    `orm:"column(lung_problem_details);size(255);null"`
	Neurological                           int       `orm:"column(neurological);null"`
	MuscularSkeletal                       int       `orm:"column(muscular_skeletal);null"`
	MentalHealth5y                         int       `orm:"column(mental_health_5y);null"`
	NeurologicalProblemDetails             string    `orm:"column(neurological_problem_details);size(255);null"`
	MuscularSkeletalProblemDetails         string    `orm:"column(muscular_skeletal_problem_details);size(255);null"`
	MuscularSkeletalNeckBack               int       `orm:"column(muscular_skeletal_neck_back);null"`
	MuscleInjuryStrainSymptoms2y           int       `orm:"column(muscle_injury_strain_symptoms_2y);null"`
	FractureBackNeckSkull                  int       `orm:"column(fracture_back_neck_skull);null"`
	FractureAnyProblems2y                  int       `orm:"column(fracture_any_problems_2y);null"`
	ArthritisDetails                       string    `orm:"column(arthritis_details);size(255);null"`
	GoutTendonitisTenosynovitis2y          int       `orm:"column(gout_tendonitis_tenosynovitis_2y);null"`
	MuscularSkeletalSportsInjury           int       `orm:"column(muscular_skeletal_sports_injury);null"`
	MuscularSkeletalJointNeckBack          int       `orm:"column(muscular_skeletal_joint_neck_back);null"`
	MentalHealthDetails                    string    `orm:"column(mental_health_details);size(255);null"`
	Alcohol28Week                          int       `orm:"column(alcohol_28_week);null"`
	IllegalDrugs5y                         int       `orm:"column(illegal_drugs_5y);null"`
	HivPositive                            int       `orm:"column(hiv_positive);null"`
	AlcoholProfessionalHelp                int       `orm:"column(alcohol_professional_help);null"`
	AlcoholDetoxHospital                   int       `orm:"column(alcohol_detox_hospital);null"`
	IllegalDrugUseDetails                  string    `orm:"column(illegal_drug_use_details);size(255);null"`
	OtherMedicalConditions                 string    `orm:"column(other_medical_conditions);size(255);null"`
	FamilyMemberCancer60                   int       `orm:"column(family_member_cancer_60);null"`
	SeekingMedicalAdviceCurrentlyDetails   string    `orm:"column(seeking_medical_advice_currently_details);size(255);null"`
	ThyroidProblem2y                       int       `orm:"column(thyroid_problem_2y);null"`
	ThryoidSurgeryRadio6m                  int       `orm:"column(thryoid_surgery_radio_6m);null"`
	CurrentMedicalTreatmentProfessions     string    `orm:"column(current_medical_treatment_professions);size(255);null"`
	FamilyMemberIllnessDetails             string    `orm:"column(family_member_illness_details);size(255);null"`
	Family2OrMoreHeartDisease              int       `orm:"column(family_2_or_more_heart_disease);null"`
	Family2OrMoreStroke                    int       `orm:"column(family_2_or_more_stroke);null"`
	KidneyDiseasePolycystic                int       `orm:"column(kidney_disease_polycystic);null"`
	Family2OrMoreDiabetes                  int       `orm:"column(family_2_or_more_diabetes);null"`
	CancerTypeDiagnosed                    string    `orm:"column(cancer_type_diagnosed);size(255);null"`
	FamilyCountColonRectal60               int       `orm:"column(family_count_colon_rectal_60);null"`
	ColonRectalBefore50                    int       `orm:"column(colon_rectal_before_50);null"`
	FamilyCountProstateCancer60            int       `orm:"column(family_count_prostate_cancer_60);null"`
	ProstateCancerBefore50                 int       `orm:"column(prostate_cancer_before_50);null"`
	Family2MoreSameCancer60                int       `orm:"column(family_2_more_same_cancer_60);null"`
	CancerBefore50                         int       `orm:"column(cancer_before_50);null"`
	FamilyCountTesticularCancer            int       `orm:"column(family_count_testicular_cancer_);null"`
	TesticularCancerBefore50               int       `orm:"column(testicular_cancer_before_50);null"`
	FamilyCountMs60                        int       `orm:"column(family_count_ms_60);null"`
	FamilyCountParkinsons60                int       `orm:"column(family_count_parkinsons_60);null"`
	FamilyCountMotorNeurone60              int       `orm:"column(family_count_motor_neurone_60);null"`
	RiskyOccupation                        string    `orm:"column(risky_occupation);size(255);null"`
	RecreationalActivities                 string    `orm:"column(recreational_activities);size(255);null"`
	TypeOfCancerWhenDiagnosed              string    `orm:"column(type_of_cancer_when_diagnosed);size(255);null"`
	CancerMedicatedTreated                 string    `orm:"column(cancer_medicated_treated);size(255);null"`
	CancerStateRemission                   string    `orm:"column(cancer_state_remission);size(255);null"`
	DiabetesFirstDiagnosed                 string    `orm:"column(diabetes_first_diagnosed);size(255);null"`
	DiabetesMedicationControl              string    `orm:"column(diabetes_medication_control);size(255);null"`
	DiabetesTests12m                       string    `orm:"column(diabetes_tests_12m);size(255);null"`
	BloodDisorderDiagnosed                 string    `orm:"column(blood_disorder_diagnosed);size(255);null"`
	BloodDisorderMedication                string    `orm:"column(blood_disorder_medication);size(255);null"`
	BloodDisorderTests12m                  string    `orm:"column(blood_disorder_tests_12m);size(255);null"`
	BloodPressureHighDiagnosed             string    `orm:"column(blood_pressure_high_diagnosed);size(255);null"`
	BloodPressureMedication                string    `orm:"column(blood_pressure_medication);size(255);null"`
	BloodPressureTest                      string    `orm:"column(blood_pressure_test);size(255);null"`
	CholesterolDiagnosed                   string    `orm:"column(cholesterol_diagnosed);size(255);null"`
	CholestrolMedication                   string    `orm:"column(cholestrol_medication);size(255);null"`
	CholesterolTest                        string    `orm:"column(cholesterol_test);size(255);null"`
	HeartConditionDiagnosed                string    `orm:"column(heart_condition_diagnosed);size(255);null"`
	HeartConditionMedication               string    `orm:"column(heart_condition_medication);size(255);null"`
	HeartConditionTest12m                  string    `orm:"column(heart_condition_test_12m);size(255);null"`
	GastroIntestinalDiagnosed              string    `orm:"column(gastro_intestinal_diagnosed);size(255);null"`
	GastroIntestinalMedication             string    `orm:"column(gastro_intestinal_medication);size(255);null"`
	GastroIntestinalCurrentState           string    `orm:"column(gastro_intestinal_current_state);size(255);null"`
	KidneyBladderDiagnosed                 string    `orm:"column(kidney_bladder_diagnosed);size(255);null"`
	KidneyBladderMedication                string    `orm:"column(kidney_bladder_medication);size(255);null"`
	KidneyBladderCurrentState              string    `orm:"column(kidney_bladder_current_state);size(255);null"`
	LungProblemDiagnosed                   string    `orm:"column(lung_problem_diagnosed);size(255);null"`
	LungProblemMedication                  string    `orm:"column(lung_problem_medication);size(255);null"`
	LungProblemCurrentState                string    `orm:"column(lung_problem_current_state);size(255);null"`
	NeurologicalDiagnosed                  string    `orm:"column(neurological_diagnosed);size(255);null"`
	NeurologicalMedication                 string    `orm:"column(neurological_medication);size(255);null"`
	NeurologicalCurrentState               string    `orm:"column(neurological_current_state);size(255);null"`
	MuscularSkeletalDiagnosed              string    `orm:"column(muscular_skeletal_diagnosed);size(255);null"`
	MuscularSkeletalMedication             string    `orm:"column(muscular_skeletal_medication);size(255);null"`
	MuscularSkeletalCurrentStatus          string    `orm:"column(muscular_skeletal_current_status);size(255);null"`
	MentalHealthDiagnosed                  string    `orm:"column(mental_health_diagnosed);size(255);null"`
	MentalHealthCausedEventIllness         string    `orm:"column(mental_health_caused_event_illness);size(255);null"`
	MentalHealthMedication                 string    `orm:"column(mental_health_medication);size(255);null"`
	AlcoholTypicalUse                      string    `orm:"column(alcohol_typical_use);size(255);null"`
	AlcoholConsultedProfessional           string    `orm:"column(alcohol_consulted_professional);size(255);null"`
	AlcoholAlcoholismOrCounselling         string    `orm:"column(alcohol_alcoholism_or_counselling);size(255);null"`
	IllegalDrugsUse                        string    `orm:"column(illegal_drugs_use);size(255);null"`
	IllegalDrugsConsultedProfessional      string    `orm:"column(illegal_drugs_consulted_professional);size(255);null"`
	IllegalDrugsHaveStoppedPeriod          string    `orm:"column(illegal_drugs_have_stopped_period);size(255);null"`
	CurrentMedicalAdviceSymptoms           string    `orm:"column(current_medical_advice_symptoms);size(255);null"`
	CurrentMedicalAdviceProfession         string    `orm:"column(current_medical_advice_profession);size(255);null"`
	CurrentMedicalAdviceTests              string    `orm:"column(current_medical_advice_tests);size(255);null"`
	FamilyMembersIllnessList               string    `orm:"column(family_members_illness_list);size(1024);null"`
	FamilyMembersMedicalTestsYouHadRelated string    `orm:"column(family_members_medical_tests_you_had_related);size(255);null"`
	CreatedAt                              time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt                              time.Time `orm:"column(updated_at);type(timestamp);null"`
	DeletedAt                              time.Time `orm:"column(deleted_at);type(timestamp);null"`
	UserId                                 int       `orm:"column(user_id);null"`
}

func (a *LifeInsurance) TableName() string {
	return "life_insurance_profile"
}

// AddLifeInsurance insert a new LifeInsurance into database and returns
// last inserted Id on success.
func AddLifeInsurance(m *LifeInsurance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLifeInsuranceById retrieves LifeInsurance by Id. Returns error if
// Id doesn't exist
func GetLifeInsuranceById(id int) (v *LifeInsurance, err error) {
	o := orm.NewOrm()
	v = &LifeInsurance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetLifeInsuranceById retrieves LifeInsurance by Id. Returns error if
// Id doesn't exist
func GetLifeInsuranceByProfileId(id int) (v *LifeInsurance, err error) {
	o := orm.NewOrm()
	v = &LifeInsurance{UserId: id}
	if err = o.Read(v, "UserId"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLifeInsurance retrieves all LifeInsurance matches certain condition. Returns empty list if
// no records exist
func GetAllLifeInsurance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LifeInsurance))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []LifeInsurance
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLifeInsurance updates LifeInsurance by Id and returns error if
// the record to be updated doesn't exist
func UpdateLifeInsuranceById(m *LifeInsurance) (err error) {
	o := orm.NewOrm()
	v := LifeInsurance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLifeInsurance deletes LifeInsurance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLifeInsurance(id int) (err error) {
	o := orm.NewOrm()
	v := LifeInsurance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LifeInsurance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type FuneralInsurances struct {
	Id              int       `orm:"column(id);auto"`
	Gender          string    `orm:"column(gender);size(255);null"`
	Age             int       `orm:"column(age);null"`
	Smoker          int       `orm:"column(smoker);null"`
	Height          int       `orm:"column(height);null"`
	Weight          int       `orm:"column(weight);null"`
	CoverAmount     int       `orm:"column(cover_amount);null"`
	ResidentCountry string    `orm:"column(resident_country);size(255);null"`
	ResidencyStatus string    `orm:"column(residency_status);size(255);null"`
	Dob             time.Time `orm:"column(dob);type(timestamp);null"`
	Created         time.Time `orm:"column(created);type(timestamp);null"`
	Updated         time.Time `orm:"column(updated);type(timestamp);null"`
	UserId          int       `orm:"column(user_id);null"`
}

func (t *FuneralInsurances) TableName() string {
	return "funeral_insurances"
}

// AddFuneralInsurances insert a new FuneralInsurances into database and returns
// last inserted Id on success.
func AddFuneralInsurances(m *FuneralInsurances) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFuneralInsurancesById retrieves FuneralInsurances by Id. Returns error if
// Id doesn't exist
func GetFuneralInsurancesById(id int) (v *FuneralInsurances, err error) {
	o := orm.NewOrm()
	v = &FuneralInsurances{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFuneralInsurances retrieves all FuneralInsurances matches certain condition. Returns empty list if
// no records exist
func GetAllFuneralInsurances(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FuneralInsurances))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []FuneralInsurances
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateFuneralInsurances updates FuneralInsurances by Id and returns error if
// the record to be updated doesn't exist
func UpdateFuneralInsurancesById(m *FuneralInsurances) (err error) {
	o := orm.NewOrm()
	v := FuneralInsurances{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFuneralInsurances deletes FuneralInsurances by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFuneralInsurances(id int) (err error) {
	o := orm.NewOrm()
	v := FuneralInsurances{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FuneralInsurances{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type HomeInsurances struct {
	Id                        int       `orm:"column(id);auto"`
	Address                   string    `orm:"column(address);size(255)"`
	City                      string    `orm:"column(city);size(45);null"`
	Postcode                  int       `orm:"column(postcode)"`
	ProposedStartDate         time.Time `orm:"column(proposed_start_date);type(date)"`
	EstimatedYearConstruction int       `orm:"column(estimated_year_construction)"`
	NumberStoreys             int       `orm:"column(number_storeys)"`
	NumberBathrooms           int       `orm:"column(number_bathrooms)"`
	NumberBedrooms            int       `orm:"column(number_bedrooms)"`
	BedroomSize               string    `orm:"column(bedroom_size);size(24)"`
	HomeDesc                  string    `orm:"column(home_desc);size(255)"`
	ExteriorWalls             string    `orm:"column(exterior_walls);size(32)"`
	RoofMaterial              string    `orm:"column(roof_material);size(32)"`
	HomeOccupied              string    `orm:"column(home_occupied);size(32)"`
	StrataPlan                string    `orm:"column(strata_plan);size(12)"`
	Pool                      string    `orm:"column(pool);size(32)"`
	Name                      string    `orm:"column(name);size(100)"`
	DobOldestPolicyHolder     time.Time `orm:"column(dob_oldest_policy_holder);type(date)"`
	EntitledNoClaim           int8      `orm:"column(entitled_no_claim)"`
	AnyBelowGround            int8      `orm:"column(any_below_ground)"`
	SecurityDevices           string    `orm:"column(security_devices);size(255)"`
	DoorSecurityDevices       string    `orm:"column(door_security_devices);size(100)"`
	WindowSecurityDevices     string    `orm:"column(window_security_devices);size(100)"`
	Alarm                     int8      `orm:"column(alarm)"`
	LandLargerNormal          int8      `orm:"column(land_larger_normal)"`
	HadPolicyCancelled        int8      `orm:"column(had_policy_cancelled)"`
	NumTimesClaimDeclined     int       `orm:"column(num_times_claim_declined)"`
	NumClaimsTotal            int       `orm:"column(num_claims_total)"`
	NumConvictions            int       `orm:"column(num_convictions)"`
	CoverStormDmgGatesFences  int8      `orm:"column(cover_storm_dmg_gates_fences)"`
	CoverAccidentalDmg        int8      `orm:"column(cover_accidental_dmg)"`
	Over50Discount            int8      `orm:"column(over_50_discount)"`
	Carports                  int       `orm:"column(carports)"`
	BalconiesDecks            int       `orm:"column(balconies_decks)"`
	Verandahs                 int       `orm:"column(verandahs)"`
	EuroKitchenAppliances     int8      `orm:"column(euro_kitchen_appliances)"`
	GraniteMarbleTiling       int8      `orm:"column(granite_marble_tiling)"`
	LargeGlazedAreas          int8      `orm:"column(large_glazed_areas)"`
	PlantationShutters        int8      `orm:"column(plantation_shutters)"`
	CurvedWalls               int8      `orm:"column(curved_walls)"`
	DuctedAircon              int8      `orm:"column(ducted_aircon)"`
	TennisCourt               int8      `orm:"column(tennis_court)"`
	IngroundPool              int8      `orm:"column(inground_pool)"`
	Watertight                int8      `orm:"column(watertight)"`
	NewHomeUnderConstruction  int8      `orm:"column(new_home_under_construction)"`
	UnderRenovation           int8      `orm:"column(under_renovation)"`
	HomeUse                   string    `orm:"column(home_use);size(100)"`
	Mortgate                  int8      `orm:"column(mortgate)"`
	InsureNameOrCompany       string    `orm:"column(insure_name_or_company);size(50)"`
	NumOwnersNamedPolicy      int       `orm:"column(num_owners_named_policy)"`
	Comments                  string    `orm:"column(comments);size(255)"`
	Status                    int       `orm:"column(status)"`
	Open                      int       `orm:"column(open)"`
	UserId                    int       `orm:"column(user_id)"`
	CreatedAt                 time.Time `orm:"column(created_at);type(timestamp)"`
	UpdatedAt                 time.Time `orm:"column(updated_at);type(timestamp)"`
}

func (t *HomeInsurances) TableName() string {
	return "home_insurances"
}

// AddHomeInsurances insert a new HomeInsurances into database and returns
// last inserted Id on success.
func AddHomeInsurances(m *HomeInsurances) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHomeInsurancesById retrieves HomeInsurances by Id. Returns error if
// Id doesn't exist
func GetHomeInsurancesById(id int) (v *HomeInsurances, err error) {
	o := orm.NewOrm()
	v = &HomeInsurances{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHomeInsurances retrieves all HomeInsurances matches certain condition. Returns empty list if
// no records exist
func GetAllHomeInsurances(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HomeInsurances))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []HomeInsurances
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateHomeInsurances updates HomeInsurances by Id and returns error if
// the record to be updated doesn't exist
func UpdateHomeInsurancesById(m *HomeInsurances) (err error) {
	o := orm.NewOrm()
	v := HomeInsurances{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHomeInsurances deletes HomeInsurances by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHomeInsurances(id int) (err error) {
	o := orm.NewOrm()
	v := HomeInsurances{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HomeInsurances{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type MortgageInsurances struct {
	Id                                  int       `orm:"column(id);auto"`
	BatchId                             int       `orm:"column(batch_id);null"`
	Gender                              string    `orm:"column(gender);size(8);null"`
	Smoker                              int       `orm:"column(smoker);null"`
	Age                                 int       `orm:"column(age);null"`
	MortgageRecent                      int       `orm:"column(mortgage_recent);null"`
	MortgageAmount                      int       `orm:"column(mortgage_amount);null"`
	LenderName                          string    `orm:"column(lender_name);size(45);null"`
	LenderBranch                        string    `orm:"column(lender_branch);size(45);null"`
	LenderCity                          string    `orm:"column(lender_city);size(45);null"`
	CountryResident                     string    `orm:"column(country_resident);size(45);null"`
	ResidentStatus                      string    `orm:"column(resident_status);size(45);null"`
	Dob                                 time.Time `orm:"column(dob);type(date);null"`
	Height                              int       `orm:"column(height);null"`
	Weight                              int       `orm:"column(weight);null"`
	CancerStatus                        int       `orm:"column(cancer_status);null"`
	CancerTypes                         string    `orm:"column(cancer_types);size(255);null"`
	DiabetesStatus                      int       `orm:"column(diabetes_status);null"`
	BloodDisorderStatus                 int       `orm:"column(blood_disorder_status);null"`
	DiabetiesComplications              int       `orm:"column(diabeties_complications);null"`
	DiabetesComplicationsDetails        string    `orm:"column(diabetes_complications_details);size(255);null"`
	HighBloodPressure                   int       `orm:"column(high_blood_pressure);null"`
	HighCholesterol                     int       `orm:"column(high_cholesterol);null"`
	HeartProblems                       int       `orm:"column(heart_problems);null"`
	HighBloodPressureTreatment          string    `orm:"column(high_blood_pressure_treatment);size(45);null"`
	HighBloodPressureHeartKidneyProblem int       `orm:"column(high_blood_pressure_heart_kidney_problem);null"`
	HighCholesterolTreatment            string    `orm:"column(high_cholesterol_treatment);size(255);null"`
	CholesterolChecked12m               int       `orm:"column(cholesterol_checked_12m);null"`
	LastCholesterolCheckStatus          string    `orm:"column(last_cholesterol_check_status);size(45);null"`
	LastCholesterolCheckReading         string    `orm:"column(last_cholesterol_check_reading);size(45);null"`
	HeartProblemsDetails                string    `orm:"column(heart_problems_details);size(255);null"`
	GastroProblemsStatus                string    `orm:"column(gastro_problems_status);size(255);null"`
	KidneyBladderProblemsStatus         string    `orm:"column(kidney_bladder_problems_status);size(255);null"`
	BreathingLungProblemStatus          string    `orm:"column(breathing_lung_problem_status);size(255);null"`
	GastroProblemsDetails               string    `orm:"column(gastro_problems_details);size(255);null"`
	HerniaRepaired                      int       `orm:"column(hernia_repaired);null"`
	NumGastritis12m                     string    `orm:"column(num_gastritis_12m);size(8);null"`
	GallstonesRemoved                   int       `orm:"column(gallstones_removed);null"`
	GallstonesOngoingIssues             int       `orm:"column(gallstones_ongoing_issues);null"`
	UlcerHospitlisation2d               int       `orm:"column(ulcer_hospitlisation_2d);null"`
	OtherGastroProblems12m              int       `orm:"column(other_gastro_problems_12m);null"`
	NumGastroIssues5y                   int       `orm:"column(num_gastro_issues_5y);null"`
	KidneyProblemDetails                string    `orm:"column(kidney_problem_details);size(255);null"`
	BreathingProblemsDetails            string    `orm:"column(breathing_problems_details);size(255);null"`
	RiskyOccupation                     string    `orm:"column(risky_occupation);size(255);null"`
	RecreationalActivities              string    `orm:"column(recreational_activities);size(255);null"`
	DiabetesMedicationControl           string    `orm:"column(diabetes_medication_control);size(255);null"`
	TypeOfCancerWhenDiagnosed           string    `orm:"column(type_of_cancer_when_diagnosed);size(255);null"`
	CancerMedicatedTreated              string    `orm:"column(cancer_medicated_treated);size(255);null"`
	CancerStateRemission                string    `orm:"column(cancer_state_remission);size(255);null"`
	DiabetesFirstDiagnosed              string    `orm:"column(diabetes_first_diagnosed);size(255);null"`
	DiabetesTests12m                    string    `orm:"column(diabetes_tests_12m);size(255);null"`
	BloodDisorderDiagnosed              string    `orm:"column(blood_disorder_diagnosed);size(255);null"`
	BloodDisorderMedication             string    `orm:"column(blood_disorder_medication);size(255);null"`
	BloodDisorderTests12m               string    `orm:"column(blood_disorder_tests_12m);size(255);null"`
	BloodPressureHighDiagnosed          string    `orm:"column(blood_pressure_high_diagnosed);size(255);null"`
	BloodPressureMedication             string    `orm:"column(blood_pressure_medication);size(255);null"`
	BloodPressureTest                   string    `orm:"column(blood_pressure_test);size(255);null"`
	CholesterolDiagnosed                string    `orm:"column(cholesterol_diagnosed);size(255);null"`
	CholestrolMedication                string    `orm:"column(cholestrol_medication);size(255);null"`
	CholesterolTest                     string    `orm:"column(cholesterol_test);size(255);null"`
	HeartConditionDiagnosed             string    `orm:"column(heart_condition_diagnosed);size(255);null"`
	HeartConditionMedication            string    `orm:"column(heart_condition_medication);size(255);null"`
	HeartConditionTest12m               string    `orm:"column(heart_condition_test_12m);size(255);null"`
	GastroIntestinalDiagnosed           string    `orm:"column(gastro_intestinal_diagnosed);size(255);null"`
	GastroIntestinalMedication          string    `orm:"column(gastro_intestinal_medication);size(255);null"`
	GastroIntestinalCurrentState        string    `orm:"column(gastro_intestinal_current_state);size(255);null"`
	KidneyBladderDiagnosed              string    `orm:"column(kidney_bladder_diagnosed);size(255);null"`
	KidneyBladderMedication             string    `orm:"column(kidney_bladder_medication);size(255);null"`
	KidneyBladderCurrentState           string    `orm:"column(kidney_bladder_current_state);size(255);null"`
	LungProblemDiagnosed                string    `orm:"column(lung_problem_diagnosed);size(255);null"`
	LungProblemMedication               string    `orm:"column(lung_problem_medication);size(255);null"`
	LungProblemCurrentState             string    `orm:"column(lung_problem_current_state);size(255);null"`
	NeurologicalDiagnosed               string    `orm:"column(neurological_diagnosed);size(255);null"`
	NeurologicalMedication              string    `orm:"column(neurological_medication);size(255);null"`
	NeurologicalCurrentState            string    `orm:"column(neurological_current_state);size(255);null"`
	MuscularSkeletalDiagnosed           string    `orm:"column(muscular_skeletal_diagnosed);size(255);null"`
	MuscularSkeletalMedication          string    `orm:"column(muscular_skeletal_medication);size(255);null"`
	MuscularSkeletalCurrentStatus       string    `orm:"column(muscular_skeletal_current_status);size(255);null"`
	MentalHealthDiagnosed               string    `orm:"column(mental_health_diagnosed);size(255);null"`
	MentalHealthCausedEventIllness      string    `orm:"column(mental_health_caused_event_illness);size(255);null"`
	MentalHealthMedication              string    `orm:"column(mental_health_medication);size(255);null"`
	AlcoholTypicalUse                   string    `orm:"column(alcohol_typical_use);size(255);null"`
	AlcoholConsultedProfessional        string    `orm:"column(alcohol_consulted_professional);size(255);null"`
	AlcoholAlcoholismOrCounselling      string    `orm:"column(alcohol_alcoholism_or_counselling);size(255);null"`
	IllegalDrugsUse                     string    `orm:"column(illegal_drugs_use);size(255);null"`
	IllegalDrugsConsultedProfessional   string    `orm:"column(illegal_drugs_consulted_professional);size(255);null"`
	IllegalDrugsHaveStoppedPeriod       string    `orm:"column(illegal_drugs_have_stopped_period);size(255);null"`
	CurrentMedicalAdviceSymptoms        string    `orm:"column(current_medical_advice_symptoms);size(255);null"`
	CurrentMedicalAdviceProfession      string    `orm:"column(current_medical_advice_profession);size(255);null"`
	CurrentMedicalAdviceTests           string    `orm:"column(current_medical_advice_tests);size(255);null"`
	FamilyMemberIllnessDetails          string    `orm:"column(family_member_illness_details);size(255);null"`
	Family2OrMoreHeartDisease           int       `orm:"column(family_2_or_more_heart_disease);null"`
	Family2OrMoreStroke                 int       `orm:"column(family_2_or_more_stroke);null"`
	KidneyDiseasePolycystic             int       `orm:"column(kidney_disease_polycystic);null"`
	Family2OrMoreDiabetes               int       `orm:"column(family_2_or_more_diabetes);null"`
	CancerTypeDiagnosed                 string    `orm:"column(cancer_type_diagnosed);size(255);null"`
	FamilyCountColonRectal60            int       `orm:"column(family_count_colon_rectal_60);null"`
	ColonRectalBefore50                 int       `orm:"column(colon_rectal_before_50);null"`
	FamilyCountProstateCancer60         int       `orm:"column(family_count_prostate_cancer_60);null"`
	ProstateCancerBefore50              int       `orm:"column(prostate_cancer_before_50);null"`
	Family2MoreSameCancer60             int       `orm:"column(family_2_more_same_cancer_60);null"`
	CancerBefore50                      int       `orm:"column(cancer_before_50);null"`
	FamilyCountTesticularCancer60       int       `orm:"column(family_count_testicular_cancer_60);null"`
	TesticularCancerBefore50            int       `orm:"column(testicular_cancer_before_50);null"`
	FamilyCountMs60                     int       `orm:"column(family_count_ms_60);null"`
	FamilyCountParkinsons60             int       `orm:"column(family_count_parkinsons_60);null"`
	FamilyCountMotorNeurone60           int       `orm:"column(family_count_motor_neurone_60);null"`
}

func (t *MortgageInsurances) TableName() string {
	return "mortgage_insurances"
}

// AddMortgageInsurances insert a new MortgageInsurances into database and returns
// last inserted Id on success.
func AddMortgageInsurances(m *MortgageInsurances) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMortgageInsurancesById retrieves MortgageInsurances by Id. Returns error if
// Id doesn't exist
func GetMortgageInsurancesById(id int) (v *MortgageInsurances, err error) {
	o := orm.NewOrm()
	v = &MortgageInsurances{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMortgageInsurances retrieves all MortgageInsurances matches certain condition. Returns empty list if
// no records exist
func GetAllMortgageInsurances(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MortgageInsurances))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MortgageInsurances
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMortgageInsurances updates MortgageInsurances by Id and returns error if
// the record to be updated doesn't exist
func UpdateMortgageInsurancesById(m *MortgageInsurances) (err error) {
	o := orm.NewOrm()
	v := MortgageInsurances{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMortgageInsurances deletes MortgageInsurances by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMortgageInsurances(id int) (err error) {
	o := orm.NewOrm()
	v := MortgageInsurances{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MortgageInsurances{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type Providers struct {
	Id             int       `orm:"column(ID);auto"`
	CompanyName    string    `orm:"column(CompanyName);size(45);null"`
	CompanyCode    string    `orm:"column(CompanyCode);size(45);null"`
	Phone1         string    `orm:"column(Phone1);size(45);null"`
	Phone2         string    `orm:"column(Phone2);size(45);null"`
	Email          string    `orm:"column(Email);size(45);null"`
	Fax            string    `orm:"column(Fax);size(45);null"`
	Website        string    `orm:"column(Website);size(45);null"`
	Information    string    `orm:"column(Information);size(255);null"`
	SectorCoverage string    `orm:"column(SectorCoverage);size(255);null"`
	Active         int       `orm:"column(active);null"`
	Country        string    `orm:"column(Country);size(45);null"`
	MailAddress    string    `orm:"column(MailAddress);size(255);null"`
	OperationHours string    `orm:"column(OperationHours);size(255);null"`
	OfficeAddress  string    `orm:"column(OfficeAddress);size(255);null"`
	Phone3         string    `orm:"column(Phone3);size(45);null"`
	MemberSince    time.Time `orm:"column(MemberSince);type(date);null"`
	Logourl        string    `orm:"column(Logourl);size(255);null"`
	Headerurl      string    `orm:"column(Headerurl);size(255);null"`
}

func (t *Providers) TableName() string {
	return "providers"
}

// AddProviders insert a new Providers into database and returns
// last inserted Id on success.
func AddProviders(m *Providers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProvidersById retrieves Providers by Id. Returns error if
// Id doesn't exist
func GetProvidersById(id int) (v *Providers, err error) {
	o := orm.NewOrm()
	v = &Providers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProviders retrieves all Providers matches certain condition. Returns empty list if
// no records exist
func GetAllProviders(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Providers))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Providers
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProviders updates Providers by Id and returns error if
// the record to be updated doesn't exist
func UpdateProvidersById(m *Providers) (err error) {
	o := orm.NewOrm()
	v := Providers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProviders deletes Providers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProviders(id int) (err error) {
	o := orm.NewOrm()
	v := Providers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Providers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type Products struct {
	Id          int       `orm:"column(id);auto"`
	CreatedAt   time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(timestamp);null"`
	DeletedAt   time.Time `orm:"column(deleted_at);type(timestamp);null"`
	Name        string    `orm:"column(name);size(255);null"`
	Description string    `orm:"column(description);size(255);null"`
	ProviderId  string    `orm:"column(provider_id);"`
	Category    string    `orm:"column(category);size(255);null"`
	Active      string    `orm:"column(active);"`
}

func (t *Products) TableName() string {
	return "products"
}

// AddProducts insert a new Products into database and returns
// last inserted Id on success.
func AddProducts(m *Products) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductsById retrieves Products by Id. Returns error if
// Id doesn't exist
func GetProductsById(id int) (v *Products, err error) {
	o := orm.NewOrm()
	v = &Products{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProducts retrieves all Products matches certain condition. Returns empty list if
// no records exist
func GetAllProducts(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Products))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Products
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateProducts updates Products by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductsById(m *Products) (err error) {
	o := orm.NewOrm()
	v := Products{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProducts deletes Products by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProducts(id int) (err error) {
	o := orm.NewOrm()
	v := Products{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Products{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type QuoteMessages struct {
	Id              int       `orm:"column(id);auto"`
	QuoteId         int       `orm:"column(quote_id);null"`
	ProviderId      int       `orm:"column(provider_id);null"`
	ProviderName    string    `orm:"column(provider_name);size(90);null"`
	Title           string    `orm:"column(title);size(255);null"`
	Body            string    `orm:"column(body);null"`
	Isread          int       `orm:"column(isread);null"`
	CreatedDatetime time.Time `orm:"column(created_datetime);type(datetime);null"`
	ReadDatetime    time.Time `orm:"column(read_datetime);type(datetime);null"`
	ParentId        int       `orm:"column(parent_id);null"`
	UserId          int       `orm:"column(user_id);null"`
}

func (t *QuoteMessages) TableName() string {
	return "quote_messages"
}

// AddQuoteMessages insert a new QuoteMessages into database and returns
// last inserted Id on success.
func AddQuoteMessages(m *QuoteMessages) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQuoteMessagesById retrieves QuoteMessages by Id. Returns error if
// Id doesn't exist
func GetQuoteMessagesById(id int) (v *QuoteMessages, err error) {
	o := orm.NewOrm()
	v = &QuoteMessages{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQuoteMessages retrieves all QuoteMessages matches certain condition. Returns empty list if
// no records exist
func GetAllQuoteMessages(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(QuoteMessages))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []QuoteMessages
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateQuoteMessages updates QuoteMessages by Id and returns error if
// the record to be updated doesn't exist
func UpdateQuoteMessagesById(m *QuoteMessages) (err error) {
	o := orm.NewOrm()
	v := QuoteMessages{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQuoteMessages deletes QuoteMessages by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQuoteMessages(id int) (err error) {
	o := orm.NewOrm()
	v := QuoteMessages{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&QuoteMessages{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

type QuoteContactRequests struct {
	Id                      int       `orm:"column(id);pk"`
	UserId                  int       `orm:"column(user_id);null"`
	QuoteType               string    `orm:"column(quote_type);size(45);null"`
	QuoteId                 int       `orm:"column(quote_id);null"`
	ProviderId              int       `orm:"column(provider_id);null"`
	ClientRequestDateTime   time.Time `orm:"column(client_request_date_time);type(datetime);null"`
	ProviderConfirmDateTime time.Time `orm:"column(provider_confirm_date_time);type(datetime);null"`
	CreatedAt               time.Time `orm:"column(created_at);type(datetime);null"`
	LastUpdated             time.Time `orm:"column(last_updated);type(datetime);null"`
	Open                    int       `orm:"column(open);null"`
}

func (t *QuoteContactRequests) TableName() string {
	return "quote_contact_requests"
}

// AddQuoteContactRequests insert a new QuoteContactRequests into database and returns
// last inserted Id on success.
func AddQuoteContactRequests(m *QuoteContactRequests) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQuoteContactRequestsById retrieves QuoteContactRequests by Id. Returns error if
// Id doesn't exist
func GetQuoteContactRequestsById(id int) (v *QuoteContactRequests, err error) {
	o := orm.NewOrm()
	v = &QuoteContactRequests{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQuoteContactRequests retrieves all QuoteContactRequests matches certain condition. Returns empty list if
// no records exist
func GetAllQuoteContactRequests(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(QuoteContactRequests))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []QuoteContactRequests
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateQuoteContactRequests updates QuoteContactRequests by Id and returns error if
// the record to be updated doesn't exist
func UpdateQuoteContactRequestsById(m *QuoteContactRequests) (err error) {
	o := orm.NewOrm()
	v := QuoteContactRequests{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQuoteContactRequests deletes QuoteContactRequests by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQuoteContactRequests(id int) (err error) {
	o := orm.NewOrm()
	v := QuoteContactRequests{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&QuoteContactRequests{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// Register models with the Beego ORM
func init() {
	orm.RegisterModel(new(LifeInsurance), new(FuneralInsurances), new(HomeInsurances), new(MortgageInsurances), new(Providers), new(Products), new(QuoteMessages), new(QuoteContactRequests), new(Profile))
}
