import React, { useState, useEffect } from 'react';
import api from '../api';
import { useNavigate, useParams } from 'react-router-dom';
import { Form, Button, Container } from 'react-bootstrap';

const EntryForm = () => {
  const [entry, setEntry] = useState({ dish: '', calories: '', fat: 0, ingredients: '' });
  const history = useNavigate();
  const { id } = useParams();

  useEffect(() => {
    if (id) {
      api.get(`/entry/${id}`)
        .then(response => setEntry(response.data))
        .catch(error => console.error(error));
    }
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;

    setEntry({ ...entry, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(entry)
    const formattedEntry = {
        ...entry,
        fat: parseFloat(entry.fat) // Convert fat to a number
      };
  
      if (id) {
        api.put(`/entry/update/${id}`, formattedEntry)
          .then(() => history('/'))
          .catch(error => console.error(error));
      } else {
        api.post('/entry/create', formattedEntry)
          .then(() => history('/'))
          .catch(error => console.error(error));
      }
  };

  return (
    <Container className="mt-4">
      <h1>{id ? 'Edit Entry' : 'Create Entry'}</h1>
      <Form onSubmit={handleSubmit}>
        <Form.Group controlId="formDish">
          <Form.Label>Dish</Form.Label>
          <Form.Control
            type="text"
            name="dish"
            value={entry.dish}
            onChange={handleChange}
            required
          />
        </Form.Group>
        <Form.Group controlId="formCalories">
          <Form.Label>Calories</Form.Label>
          <Form.Control
            type="text"
            name="calories"
            value={entry.calories}
            onChange={handleChange}
            required
          />
        </Form.Group>
        <Form.Group controlId="formFat">
          <Form.Label>Fat</Form.Label>
          <Form.Control
            type="number"
            name="fat"
            value={entry.fat}
            onChange={handleChange}
            required
          />
        </Form.Group>
        <Form.Group controlId="formIngredients">
          <Form.Label>Ingredients</Form.Label>
          <Form.Control
            type="text"
            name="ingredients"
            value={entry.ingredients}
            onChange={handleChange}
            required
          />
        </Form.Group>
        <Button variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </Container>
  );
};

export default EntryForm;
