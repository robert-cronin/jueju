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
import { Card, CardContent, Typography, Box } from "@mui/material";

interface PoemCardProps {
  title: string;
  content: string;
  translation: string;
}

const PoemCard: React.FC<PoemCardProps> = ({ title, content, translation }) => {
  return (
    <Card sx={{ minWidth: 275, maxWidth: 400, margin: "16px auto" }}>
      <CardContent>
        <Typography variant="h5" component="div" gutterBottom>
          {title}
        </Typography>
        <Box mb={2}>
          <Typography variant="body1" color="text.secondary">
            {content}
          </Typography>
        </Box>
        <Typography variant="body2">
          Translation:
          <br />
          {translation}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default PoemCard;
