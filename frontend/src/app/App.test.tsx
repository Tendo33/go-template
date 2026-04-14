import { render, screen, waitFor } from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

import { App } from "./App";

describe("App", () => {
  const fetchMock = vi.fn();

  beforeEach(() => {
    vi.stubGlobal("fetch", fetchMock);
  });

  afterEach(() => {
    vi.unstubAllGlobals();
    fetchMock.mockReset();
  });

  it("renders the template title and health status from backend", async () => {
    fetchMock.mockResolvedValue({
      ok: true,
      json: async () => ({ status: "ok", service: "go-template" })
    });

    render(<App />);

    expect(
      screen.getByRole("heading", { name: /go fullstack template/i })
    ).toBeInTheDocument();

    await waitFor(() => {
      expect(screen.getByText(/backend status: ok/i)).toBeInTheDocument();
      expect(screen.getByText(/service: go-template/i)).toBeInTheDocument();
    });
  });
});
