// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { initializeAuth, browserLocalPersistence } from "firebase/auth"; // TODO: Add SDKs for Firebase products that you want to use

// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDHREF4Gb4SygCNWLZmpXzKDL4UgoXHnMo",
  authDomain: "somev2.firebaseapp.com",
  projectId: "somev2",
  storageBucket: "somev2.appspot.com",
  messagingSenderId: "25867038734",
  appId: "1:25867038734:web:eca8df725f25fbb4800c5f",
  measurementId: "G-KW6NNP6DY0",
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = initializeAuth(app, {
  persistence: browserLocalPersistence,
});

export { auth };
