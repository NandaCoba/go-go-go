import React, { useEffect, useState } from 'react';
import axios from 'axios';

const App = () => {
  const [data, setData] = useState([]);
  const [name, setName] = useState('');
  const [age, setAge] = useState('');
  const [editing, setEditing] = useState(null);
  const api = 'http://localhost:3000/';

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = () => {
    axios
      .get(api)
      .then((result) => {
        setData(Array.isArray(result.data.data) ? result.data.data : []);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  // Create new data
  const createData = () => {
    if (!name.trim() || !age) {
      alert('Please fill all fields');
      return;
    }

    // Membuat FormData object
    const formData = new FormData();
    formData.append('Name', name);
    formData.append('Age', age);

    axios
      .post(api, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      })
      .then((result) => {
        fetchData();
        setName('');
        setAge('');
      })
      .catch((err) => {
        console.log(err);
        alert('Failed to create data');
      });
  };

  // Update data
  const updateData = () => {
    if (!name.trim() || !age) {
      alert('Please fill all fields');
      return;
    }

    // Membuat FormData object untuk update
    const formData = new FormData();
    formData.append('Name', name);
    formData.append('Age', age);

    axios
      .put(`${api}${editing.Id}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      })
      .then((result) => {
        fetchData();
        setEditing(null);
        setName('');
        setAge('');
      })
      .catch((err) => {
        console.log(err);
        alert('Failed to update data');
      });
  };

  // Delete data
  const deleteData = (id) => {
    if (window.confirm('Are you sure you want to delete this item?')) {
      axios
        .delete(`${api}${id}`)
        .then(() => {
          fetchData();
        })
        .catch((err) => {
          console.log(err);
          alert('Failed to delete data');
        });
    }
  };

  // Set form for editing
  const editData = (item) => {
    setEditing(item);
    setName(item.Name);
    setAge(item.Age);
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      <h1 className="text-3xl font-bold text-center mb-8 text-black">User Management <span className=' text-blue-400'>System</span></h1>
      
      {/* Form Section */}
      <div className="bg-white rounded-lg shadow-md p-6 mb-8">
        <div className="flex flex-col md:flex-row gap-4">
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Enter name"
            className="flex-1 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <input
            type="number"
            value={age}
            onChange={(e) => setAge(e.target.value)}
            placeholder="Enter age"
            className="flex-1 md:w-32 px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          {editing ? (
            <div className="flex gap-2">
              <button
                onClick={updateData}
                className="px-6 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors"
              >
                Update
              </button>
              <button
                onClick={() => {
                  setEditing(null);
                  setName('');
                  setAge('');
                }}
                className="px-6 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors"
              >
                Cancel
              </button>
            </div>
          ) : (
            <button
              onClick={createData}
              className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
            >
              Create
            </button>
          )}
        </div>
      </div>

      {/* Data Display Section */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {data.length === 0 ? (
          <div className="col-span-full text-center text-gray-500 py-8">
            No data available
          </div>
        ) : (
          data.map((item) => (
            <div 
              key={item.Id}
              className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow"
            >
              <h2 className="text-xl font-semibold mb-2">{item.Name || 'No Name'}</h2>
              <p className="text-gray-600 mb-2">Age: {item.Age || 'N/A'}</p>
              <p className="text-gray-500 text-sm mb-4">
                Created: {new Date(item.CreatedAt).toLocaleDateString()}
              </p>
              <div className="flex gap-2">
                <button
                  onClick={() => editData(item)}
                  className="flex-1 px-4 py-2 bg-yellow-500 text-white rounded-lg hover:bg-yellow-600 transition-colors"
                >
                  Edit
                </button>
                <button
                  onClick={() => deleteData(item.Id)}
                  className="flex-1 px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors"
                >
                  Delete
                </button>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default App;