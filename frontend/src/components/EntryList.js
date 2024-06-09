import React, { useEffect, useState } from 'react';

import { Link } from 'react-router-dom';
import { Table, Button } from 'react-bootstrap';
import api from '../api';

const EntryList = () => {
  const [entries, setEntries] = useState([]);

  useEffect(() => {
    api.get('/entries')
      .then(response => setEntries(response.data))
      .catch(error => console.error(error));
  }, []);

  const deleteEntry = (id) => {
    api.delete(`/delete/${id}`)
      .then(() => setEntries(entries.filter(entry => entry.id !== id)))
      .catch(error => console.error(error));
  };

  return (
    <div className="container mt-4">
      <h1>Entries</h1>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Dish</th>
            <th>Calories</th>
            <th>Fat</th>
            <th>Ingredients</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {entries.map(entry => (
            <tr key={entry.id}>
              <td>{entry.dish}</td>
              <td>{entry.calories}</td>
              <td>{entry.fat}</td>
              <td>{entry.ingredients}</td>
              <td>
                <Button variant="info" as={Link} to={`/entry/${entry.id}`}>View</Button>{' '}
                <Button variant="warning" as={Link} to={`/edit/${entry.id}`}>Edit</Button>{' '}
                <Button variant="danger" onClick={() => deleteEntry(entry.id)}>Delete</Button>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};

export default EntryList;
