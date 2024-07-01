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
import PoemCard from "@/components/PoemCard";

const mockPoems = [
  {
    title: "静夜思 (Jìng Yè Sī)",
    content: "床前明月光，疑是地上霜。举头望明月，低头思故乡。",
    translation:
      "Moonlight before my bed,\nCould it be frost instead?\nHead up I watch the moon,\nHead down I miss my home.",
  },
  {
    title: "春晓 (Chūn Xiǎo)",
    content: "春眠不觉晓，处处闻啼鸟。夜来风雨声，花落知多少。",
    translation:
      "Spring dawn I fail to notice,\nHear everywhere birds sing.\nLast night the sound of wind and rain,\nHow many flowers fell?",
  },
  {
    title: "登鹳雀楼 (Dēng Guàn Què Lóu)",
    content: "白日依山尽，黄河入海流。欲穷千里目，更上一层楼。",
    translation:
      "The white sun sets behind the mountains,\nThe Yellow River flows into the sea.\nTo see a thousand miles further,\nClimb one more storey of the tower.",
  },
];

const Feed: React.FC = () => {
  return (
    <Container maxWidth="md">
      <Typography variant="h4" component="h1" gutterBottom align="center">
        Poetry Feed
      </Typography>
      {mockPoems.map((poem, index) => (
        <PoemCard key={index} {...poem} />
      ))}
    </Container>
  );
};

export default Feed;
