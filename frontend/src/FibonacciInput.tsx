import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const FibonacciInput: React.FC = () => {
    const [number, setNumber] = useState<number | ''>('');
    const navigate = useNavigate();

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setNumber(Number(e.target.value));
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (number >= 0) {
            navigate(`/fibonacci/${number}`);
        } else {
            alert("Please enter a non-negative integer.");
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="number"
                value={number}
                onChange={handleInputChange}
                placeholder="Enter a non-negative integer"
                required
            />
            <button type="submit">Generate</button>
        </form>
    );
};

export default FibonacciInput;
