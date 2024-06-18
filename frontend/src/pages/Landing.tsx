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
import { Container, Button, Typography, Paper, Divider } from "@mui/material";
import { styled } from "@mui/system";
import useAuth from "@/hooks/useAuth";

const LandingContainer = styled(Container)`
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  height: 100vh;
  padding: 20px;
  @media (max-width: 600px) {
    flex-direction: column;
  }
`;

const LeftSection = styled("div")`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  min-width: 400px; 
`;

const RightSection = styled(Paper)`
  padding: 20px;
  margin-left: 20px;
  @media (max-width: 600px) {
    margin-left: 0;
    margin-top: 20px;
  }
`;

const VerticalDivider = styled(Divider)`
  height: auto;
  margin: 0 20px;
  @media (max-width: 600px) {
    width: 100%;
    height: 1px;
    margin: 20px 0;
  }
`;

const LandingPage: React.FC = () => {
  const { goToLogin } = useAuth();
  return (
    <LandingContainer>
      <LeftSection>
        <Typography variant="h2" gutterBottom>
          绝句 (jué jù)
        </Typography>
        <Typography variant="h5" gutterBottom>
          AI-powered Chinese poetry.
        </Typography>
        <Button variant="contained" color="primary" onClick={goToLogin}>
          Login (or Register)
        </Button>
      </LeftSection>
      <VerticalDivider orientation="vertical" flexItem />
      <RightSection elevation={3}>
        <Typography variant="body1" gutterBottom>
          静夜思 (Jìng Yè Sī) - Li Bai
        </Typography>
        <Typography variant="body2" gutterBottom>
          静夜思 唐 · 李白 床前明月光， 疑是地上霜。 举头望明月， 低头思故乡。
        </Typography>
        <Typography variant="body1" gutterBottom>
          Translation:
        </Typography>
        <Typography variant="body2">
          Tang Dynasty · Li Bai
          <br />
          Before my bed, the moonlight is so bright, It seems like frost upon
          the ground. I raise my head to gaze at the bright moon, And lower it,
          thinking of my hometown.
        </Typography>
      </RightSection>
    </LandingContainer>
  );
};

export default LandingPage;
