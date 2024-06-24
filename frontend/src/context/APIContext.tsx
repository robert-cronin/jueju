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

import React, { ReactNode, createContext } from "react";
import * as clients from "@clients/v1.0";

interface APIContextType {
  api: clients.DefaultApi;
}

const APIContext = createContext<APIContextType | undefined>(undefined);

interface APIProviderProps {
  children: ReactNode;
}

const APIProvider: React.FC<APIProviderProps> = ({ children }) => {
  const basePath = import.meta.env.VITE_API_BASE_PATH;
  const configuration = new clients.Configuration({
    basePath,
    baseOptions: {
      withCredentials: true,
    },
    
  });
  const api = new clients.DefaultApi(configuration);

  return <APIContext.Provider value={{ api }}>{children}</APIContext.Provider>;
};

export { APIProvider, APIContext };
export type { APIContextType };
