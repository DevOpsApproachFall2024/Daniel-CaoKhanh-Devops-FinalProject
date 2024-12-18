import React from 'react';
import { Link } from 'react-router-dom';
import FibonacciInput from './FibonacciInput';
import './WelcomePage.css';

const WelcomePage: React.FC = () => {
    return (
        <div>
            <h1>Welcome!</h1>
            <div className="button-container">
                <Link to="/tony">
                    <button>Watch Tony&apos;s Video</button>
                </Link>
                <Link to="/rick">
                    <button>Watch Rick&apos;s Video</button>
                </Link>
            </div>
            <h2>Generate Fibonacci Sequence</h2>
            <FibonacciInput />
        </div>
    );
};

export default WelcomePage;
