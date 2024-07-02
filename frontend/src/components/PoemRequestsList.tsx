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

import React, { useEffect, useState } from "react";
import { List, ListItem, ListItemText, Typography, Box } from "@mui/material";
import useAPI from "@/hooks/useAPI";
import { PoemRequest } from "@clients/v1.0";

const PoemRequestsList: React.FC = () => {
  const [poemRequests, setPoemRequests] = useState<PoemRequest[]>([]);
  const [error, setError] = useState<string | null>(null);
  const { api } = useAPI();

  useEffect(() => {
    const fetchPoemRequests = async () => {
      try {
        const response = await api.getUserPoemRequests();
        setPoemRequests(response.data);
      } catch (err) {
        setError("Failed to fetch poem requests. Please try again.");
        console.error(err);
      }
    };

    fetchPoemRequests();
  }, [api]);

  if (error) {
    return <Typography color="error">{error}</Typography>;
  }

  return (
    <Box sx={{ mt: 3 }}>
      <Typography variant="h6">Your Poem Requests:</Typography>
      <List>
        {poemRequests.map((request) => (
          <ListItem key={request.id}>
            <ListItemText
              primary={request.prompt}
              secondary={`Status: ${request.status} | Created: ${new Date(request.created_at).toLocaleString()}`}
            />
          </ListItem>
        ))}
      </List>
    </Box>
  );
};

export default PoemRequestsList;
