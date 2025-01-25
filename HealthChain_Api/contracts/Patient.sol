//// SPDX-License-Identifier: MIT
//pragma solidity ^0.8.0;
//
//contract PatientContract {
//    struct EmergencyContact {
//        string name;
//        string relationship;
//        string phoneNumber;
//    }
//
//    struct Patient {
//        uint id;
//        string name;s
//        uint age;
//        string medicalHistory;
//        string patientAddress;
//        string phoneNumber;
//        string email;
//        string allergies;
//        string currentMedications;
//        EmergencyContact emergencyContact;
//    }
//
//    mapping(uint => Patient) private patients;
//    uint private nextPatientId;
//    address private owner;
//
//    event PatientAdded(uint id, string name);
//    event PatientUpdated(uint id, string name);
//
//    constructor() {
//        owner = msg.sender;
//        nextPatientId = 1;
//    }
//
//    modifier onlyOwner() {
//        require(msg.sender == owner, "Only owner can call this function");
//        _;
//    }
//
//    function addPatient(
//        string memory name,
//        uint age,
//        string memory medicalHistory,
//        string memory patientAddress,
//        string memory phoneNumber,
//        string memory email,
//        string memory allergies,
//        string memory currentMedications,
//        string memory emergencyContactName,
//        string memory emergencyContactRelationship,
//        string memory emergencyContactPhoneNumber
//    ) public onlyOwner {
//        Patient memory newPatient = Patient({
//            id: nextPatientId,
//            name: name,
//            age: age,
//            medicalHistory: medicalHistory,
//            patientAddress: patientAddress,
//            phoneNumber: phoneNumber,
//            email: email,
//            allergies: allergies,
//            currentMedications: currentMedications,
//            emergencyContact: EmergencyContact({
//            name: emergencyContactName,
//            relationship: emergencyContactRelationship,
//            phoneNumber: emergencyContactPhoneNumber
//        })
//        });
//
//        patients[nextPatientId] = newPatient;
//        emit PatientAdded(nextPatientId, name);
//        nextPatientId++;
//    }
//
//    function getPatient(uint patientId) public view returns (Patient memory) {
//        require(patientId > 0 && patientId < nextPatientId, "Patient does not exist");
//        return patients[patientId];
//    }
//
//    function updatePatient(
//        uint patientId,
//        string memory name,
//        uint age,
//        string memory medicalHistory,
//        string memory patientAddress,
//        string memory phoneNumber,
//        string memory email,
//        string memory allergies,
//        string memory currentMedications,
//        string memory emergencyContactName,
//        string memory emergencyContactRelationship,
//        string memory emergencyContactPhoneNumber
//    ) public onlyOwner {
//        require(patientId > 0 && patientId < nextPatientId, "Patient does not exist");
//
//        Patient storage patient = patients[patientId];
//        patient.name = name;
//        patient.age = age;
//        patient.medicalHistory = medicalHistory;
//        patient.patientAddress = patientAddress;
//        patient.phoneNumber = phoneNumber;
//        patient.email = email;
//        patient.allergies = allergies;
//        patient.currentMedications = currentMedications;
//        patient.emergencyContact = EmergencyContact({
//            name: emergencyContactName,
//            relationship: emergencyContactRelationship,
//            phoneNumber: emergencyContactPhoneNumber
//        });
//
//        emit PatientUpdated(patientId, name);
//    }
//}
