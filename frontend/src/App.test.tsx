import { render } from '@testing-library/react';
import App from './App';

test('renders the welcome message', () => {
  const { getByText } = render(<App />); // Use destructured render methods directly
  const messageElement = getByText(/Hi, I am Daniel and I am single since a long time ago/i);
  expect(messageElement).toBeInTheDocument();
});
