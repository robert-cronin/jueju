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

import React, { useState } from "react";
import { Container, Typography, Box, Button, Alert } from "@mui/material";
import PoemRequestForm from "@/components/PoemRequestForm";
import PoemList from "@/components/PoemList";
import useAuth from "@/hooks/useAuth";

const Create: React.FC = () => {
  const { user } = useAuth();
  const [showAlert, setShowAlert] = useState(false);

  const handleGetMoreCredits = () => {
    setShowAlert(true);
    setTimeout(() => setShowAlert(false), 3000); // Hide alert after 3 seconds
  };

  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom align="center">
        Create a New Poem
      </Typography>
      <Box display="flex" justifyContent="center" alignItems="center" mb={2}>
        <Typography variant="h6" mr={2}>
          Poem Credits: {user?.poem_credits || 0}
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={handleGetMoreCredits}
        >
          Get More Credits
        </Button>
      </Box>
      {showAlert && (
        <Alert severity="info" onClose={() => setShowAlert(false)}>
          This feature is currently in development. Stay tuned for updates!
        </Alert>
      )}
      <PoemRequestForm />
      <Box mt={4}>
        <PoemList status="pending" />
      </Box>
      <Box mt={4}>
        <PoemList status="failed" />
      </Box>
    </Container>
  );
};

export default Create;
