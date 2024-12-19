import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import './Fibonacci.css';

const Fibonacci: React.FC = () => {
    const { number } = useParams<{ number: string }>();
    const [result, setResult] = useState<string>('');

    useEffect(() => {
        const fetchFibonacci = async () => {
            try {
                const response = await fetch(`http://localhost:8080/fibonacci/${number}`);
                if (!response.ok) throw new Error('Failed to fetch');
                const data = await response.text();
                setResult(data);
            } catch (error) {
                console.error("Error fetching Fibonacci sequence:", error);
                setResult("Error fetching the Fibonacci sequence.");
            }
        };

        if (number) {
            fetchFibonacci();
        }
    }, [number]);

    return (
        <div>
            <h1>Fibonacci Sequence for {number}</h1>
            {result && <p>{result}</p>}
        </div>
    );
};

export default Fibonacci;
