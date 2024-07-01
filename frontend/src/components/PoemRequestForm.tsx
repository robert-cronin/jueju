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
import { TextField, Button, Typography, Box } from "@mui/material";
import useApi from "@/hooks/useAPI";
import { PoemRequest } from "@clients/v1.0";

interface PoemRequestFormProps {
  onPoemCreated?: () => void;
}

const PoemRequestForm: React.FC<PoemRequestFormProps> = ({ onPoemCreated }) => {
  const [prompt, setPrompt] = useState("");
  const [poemRequest, setPoemRequest] = useState<PoemRequest | null>(null);
  const [error, setError] = useState<string | null>(null);
  const { api } = useApi();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const result = await api.requestPoem({ prompt });
      setPoemRequest(result.data);
      setError(null);
      if (onPoemCreated) {
        onPoemCreated();
      }
    } catch (err) {
      setError("Failed to request poem. Please try again.");
      console.error(err);
    }
  };

  return (
    <Box component="form" onSubmit={handleSubmit} sx={{ mt: 3 }}>
      <TextField
        fullWidth
        label="Poem Prompt"
        value={prompt}
        onChange={(e) => setPrompt(e.target.value)}
        margin="normal"
        required
      />
      <Button type="submit" variant="contained" color="primary" sx={{ mt: 2 }}>
        Request Poem
      </Button>
      {error && (
        <Typography color="error" sx={{ mt: 2 }}>
          {error}
        </Typography>
      )}
      {poemRequest && (
        <Box sx={{ mt: 3 }}>
          <Typography variant="h6">Your Poem:</Typography>
          <Typography>{poemRequest.poem}</Typography>
        </Box>
      )}
    </Box>
  );
};

export default PoemRequestForm;
