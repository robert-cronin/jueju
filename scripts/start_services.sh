#!/bin/bash

# Copyright 2024 Robert Cronin
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Start a new tmux session named 'jueju'
tmux new-session -d -s jueju

# Split the window into three panes
tmux split-window -h
tmux split-window -v

# Select the first pane and start the backend
tmux select-pane -t 0
tmux send-keys "cd backend && go run main.go" C-m

# Select the second pane and start the frontend
tmux select-pane -t 1
tmux send-keys "cd frontend && npm run dev" C-m

# Select the third pane and start the AI service
tmux select-pane -t 2
tmux send-keys "cd ai_service && python main.py" C-m

# Attach to the tmux session
tmux attach-session -t jueju
