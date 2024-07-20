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

import React, { useState, useEffect } from "react";
import { Box, Typography, Button, CircularProgress, Grid } from "@mui/material";
import PoemCard from "./PoemCard";
import useAPI from "@/hooks/useAPI";
import { PoemRequest } from "@clients/v1.0";

interface PoemListProps {
  status: "completed" | "pending" | "failed";
  showRefresh?: boolean;
}

const PoemList: React.FC<PoemListProps> = ({ status, showRefresh = true }) => {
  const [poems, setPoems] = useState<PoemRequest[]>([]);
  const [loading, setLoading] = useState(true);
  const { api } = useAPI();

  const fetchPoems = async () => {
    setLoading(true);
    try {
      const response = await api.getUserPoemRequests();
      setPoems(response.data.filter((poem) => poem.status === status));
    } catch (error) {
      console.error("Error fetching poems:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPoems();
  }, [status]);

  const handleRefresh = () => {
    fetchPoems();
  };

  const handleRetry = async (poemId: string) => {
    try {
      await api.requestPoem({
        prompt: poems.find((p) => p.id === poemId)?.prompt || "",
      });
      fetchPoems();
    } catch (error) {
      console.error("Error retrying poem:", error);
    }
  };

  return (
    <Box>
      <Box
        display="flex"
        justifyContent="space-between"
        alignItems="center"
        mb={2}
      >
        <Typography variant="h5">
          {status.charAt(0).toUpperCase() + status.slice(1)} Poems
        </Typography>
        {showRefresh && (
          <Button
            variant="contained"
            onClick={handleRefresh}
            disabled={loading}
          >
            Refresh
          </Button>
        )}
      </Box>
      {loading ? (
        <CircularProgress />
      ) : (
        <Grid container spacing={2}>
          {poems.map((poem) => (
            <Grid item xs={12} sm={6} md={4} key={poem.id}>
              <PoemCard
                title={poem.prompt}
                content={poem.poem || "Generating..."}
                translation={poem.status}
                onRetry={
                  status === "failed" ? () => handleRetry(poem.id) : undefined
                }
              />
            </Grid>
          ))}
        </Grid>
      )}
    </Box>
  );
};

export default PoemList;
