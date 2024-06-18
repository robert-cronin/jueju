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
// src/context/AuthContext.tsx

import LandingPage from "@/pages/Landing";
import LoadingScreen from "@/pages/Loading";
import React, { createContext, useState, useEffect, ReactNode } from "react";
import { useNavigate } from "react-router-dom";
import { useAPI } from "@/context/APIContext";

type AuthContextType = {
  user: object | null;
  goToLogin: () => void;
  goToLogout: () => void;
  loading: boolean;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<object | null>(null);
  const [loading, setLoading] = useState(true);

  const { api } = useAPI();

  const navigate = useNavigate();

  const goToLogin = () => {
    api.login();
  };

  const goToLogout = () => {
    api.logout();
    setUser(null);
    navigate("/");
  };

  const getUser = async () => {
    try {
      setLoading(true);
      const response = await api.getUser();
      const { data } = response;
      setUser(data);
    } catch (error) {
      setUser(null);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getUser();
  }, []);

  if (loading) {
    return <LoadingScreen />;
  }

  return (
    <AuthContext.Provider value={{ user, goToLogin, goToLogout, loading }}>
      {user ? children : <LandingPage />}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
export { AuthContext };
export type { AuthContextType };
