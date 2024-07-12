import { useState } from 'react';

function App() {
  const [id, setId] = useState('');
  const [response, setResponse] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/data?id=${id}`);
      if (!res.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await res.text();
      setResponse(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center text-black p-4 bg-gradient-to-r from-gray-500 to-sky-300">
      <h1 className="text-4xl font-bold mb-6 mt-6">Codeforces Submissions</h1>
      <form onSubmit={handleSubmit} className="w-full max-w-md  bg-gradient-to-r from-gray-300 to-gray-500 rounded-lg shadow-md p-6 space-y-4 mt-6">
        <div>
          <label htmlFor="id" className="block text-black font-bold mb-2">Enter CF ID:</label>
          <input
            type="text"
            id="id"
            value={id}
            onChange={(e) => setId(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg text-black font-[400]"
            required
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-700 font-[400]"
        >
          {loading ? 'Loading...' : 'Enter'}
        </button>
      </form>
      {error && <p className="text-red-500 mt-4">{error}</p>}
      {response && (
        <pre className="mt-6 bg-transparent text-black font-[800] p-4 rounded-lg w-full max-w-md overflow-auto">
          {response}
        </pre>
      )}
    </div>
  );
}

export default App;
