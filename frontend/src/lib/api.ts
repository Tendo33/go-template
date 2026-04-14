export type HealthResponse = {
  status: string;
  service: string;
};

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? "";

export async function fetchHealth(): Promise<HealthResponse> {
  const response = await fetch(`${API_BASE_URL}/health`);

  if (!response.ok) {
    throw new Error(`health request failed: ${response.status}`);
  }

  return response.json() as Promise<HealthResponse>;
}
