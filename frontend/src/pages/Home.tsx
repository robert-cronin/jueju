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
import PoemList from "@/components/PoemList";

const Home: React.FC = () => {
  return (
    <Container maxWidth="lg">
      <Typography variant="h2" align="center" gutterBottom>
        Welcome to JueJu
      </Typography>
      <Typography variant="body1" align="center" gutterBottom>
        Create and explore AI-generated Chinese poetry.
      </Typography>
      <PoemList status="completed" />
    </Container>
  );
};

export default Home;
