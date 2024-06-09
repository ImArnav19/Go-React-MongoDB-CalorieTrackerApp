import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Container } from 'react-bootstrap';

const EntryDetails = () => {
  const [entry, setEntry] = useState({});
  const { id } = useParams();

  useEffect(() => {
    axios.get(`/entry/${id}`)
      .then(response => setEntry(response.data))
      .catch(error => console.error(error));
  }, [id]);

  return (
    <Container className="mt-4">
      <h1>{entry.dish}</h1>
      <p><strong>ID:</strong> {entry.id}</p>
      <p><strong>Calories:</strong> {entry.calories}</p>
      <p><strong>Fat:</strong> {entry.fat}</p>
      <p><strong>Ingredients:</strong> {entry.ingredients}</p>
    </Container>
  );
};

export default EntryDetails;
