// Copyright 2024 Robert Cronin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from "react";
import { Container, Typography } from "@mui/material";
import PoemRequestForm from "@/components/PoemRequestForm";
import { useNavigate } from "react-router-dom";

const Create: React.FC = () => {
  const navigate = useNavigate();

  const handlePoemCreated = () => {
    // Redirect to home after poem is created
    navigate("/");
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" component="h1" gutterBottom align="center">
        Create a New Poem
      </Typography>
      <PoemRequestForm onPoemCreated={handlePoemCreated} />
    </Container>
  );
};

export default Create;
