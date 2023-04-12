'use strict';

const {Contract, Context} = require('fabric-contract-api');
const File = require('./File.js');

class Patient {
   
    constructor(){

    }

    static createPatient(PatientId, name, age, sex, address, email, contact, hospital, Files, Doctors, bills, insurance, Requests){

    }    

    static createFile(PatientId, fileName, doctor, date, category, comments)
    {
        // creatte a file add it to files of patient
        let file = File.newFile(fileName, doctorId, date, category, comments);
        this.Files.push(file);
    }
    getFile(PatientId, fileName){
        // read a file from patients list
        this.Files
    }


}
module.exports= Patient;