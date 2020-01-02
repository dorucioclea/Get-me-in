import React from 'react';
import ReactDOM from 'react-dom';
import {Login, unregister} from "./ts"
import './css/index.css';
import './css/App.css';

ReactDOM.render(<Login />, document.getElementById('root'));

unregister();
// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA

