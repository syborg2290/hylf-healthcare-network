'use strict';

const {Contract, Context} = require('fabric-contract-api');

const Patient = require('./Patient.js');
const PatientList = require('./PatientList.js');
const File = require('./File.js');

class HealthCareContext extends Context {

    constructor() {
        super();
        // All papers are held in a list of papers
        this.patientList = new PatientList(this);
    }

}

class Healthcare extends Contract{

    constructor(){
        super();
    }

    createContext() {
        return new HealthCareContext();
    }

    async newPatient(ctx, PatientId, name, age, sex, address, email, contact, hospital, Files, Doctors, bills, insurance, Requests){
        let patient = Patient.createPatient(PatientId, name, age, sex, address, email, contact, hospital, Files, Doctors, bills, insurance, Requests);
        await ctx.patientList.addPatient(patient);
    }
    
    async addFile(ctx, PatientId, fileName, doctorId, date, category, comments){
        //get patient details from patientlist using id add new file in the file data, update doctors list
        
        let patient = await ctx.patientList.getpatient(PatientId);
        patient.createFile( PatientId, fileName, doctorId, date, category, comments);
        await ctx.patientList.updatePatient(patient);
    }

    async requestAccess(ctx, PatientId, filename, doctorId, from, to, requestOwner){
        //org asks request to view a patients data

    }

    async approveAccess(ctx, PatientId, filename, doctorId, requestOwner){
        //hospital approves request made by an org
    }

    async viewFile(ctx, PatientId, fileName, doctorId){
        //the hospital, patient or an org with granted request can view the documents
        let patient =  await ctx.patientList.getpatient(PatientId);
        let file = patient.getFile(PatientId, fileName);
        return file;
    }


    
}