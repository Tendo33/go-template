import { useEffect, useState } from "react";

import { fetchHealth, type HealthResponse } from "../lib/api";

type HealthState =
  | { kind: "loading" }
  | { kind: "success"; data: HealthResponse }
  | { kind: "error"; message: string };

export function App() {
  const [health, setHealth] = useState<HealthState>({ kind: "loading" });

  useEffect(() => {
    let active = true;

    void fetchHealth()
      .then((data) => {
        if (active) {
          setHealth({ kind: "success", data });
        }
      })
      .catch((error: Error) => {
        if (active) {
          setHealth({ kind: "error", message: error.message });
        }
      });

    return () => {
      active = false;
    };
  }, []);

  return (
    <main className="app-shell">
      <section className="hero-card">
        <p className="eyebrow">Go + Gin + React</p>
        <h1>Go Fullstack Template</h1>
        <p className="lede">
          一个面向全栈项目起步的模板仓库，默认提供 Gin 后端、
          React/Vite 前端与基础工程化配置。
        </p>
        {health.kind === "loading" ? <p>Loading backend status...</p> : null}
        {health.kind === "error" ? <p>Backend error: {health.message}</p> : null}
        {health.kind === "success" ? (
          <div>
            <p>Backend status: {health.data.status}</p>
            <p>Service: {health.data.service}</p>
          </div>
        ) : null}
      </section>
    </main>
  );
}
