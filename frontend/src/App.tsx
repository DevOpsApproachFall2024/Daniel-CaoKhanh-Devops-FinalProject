import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import WelcomePage from './WelcomePage';
import TonyVideo from './TonyVideo';
import RickVideo from './RickVideo';
import Fibonacci from './Fibonacci';

const App: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<WelcomePage />} />
                <Route path="/tony" element={<TonyVideo />} />
                <Route path="/rick" element={<RickVideo />} />
                <Route path="/fibonacci/:number" element={<Fibonacci />} />
            </Routes>
        </Router>
    );
};

export default App;
