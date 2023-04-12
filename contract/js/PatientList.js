'use-strict';

const Patient = require('./Patient.js');

class PatientList {

    constructor(ctx) {
        super(ctx, 'org.papernet.paper');
        this.use(CommercialPaper);
    }

    async addPatient(patient) {
        return this.addState(patient);
    }

    async getPatient(patientId) {
        return this.getState(patientId);
    }

    async updatePatient(patient) {
        return this.updateState(patient);
    }
}
module.exports= PatientList;