// @vitest-environment jsdom
import "@testing-library/jest-dom/vitest";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { PostView } from "@/components/post-view";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  createMemoryHistory,
  RouterProvider,
  createRootRoute,
  createRouter,
} from "@tanstack/react-router";
import api from "@/api";
import type { Post } from "@/types";

// 1. Mock the API module
vi.mock("@/api", () => ({
  default: {
    toggleLikePost: vi.fn(),
    toggleBookmarkPost: vi.fn(),
  },
}));

// 2. Helper to wrap component with necessary Providers
const renderWithProviders = (ui: React.ReactElement) => {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });

  // Simple router mock for TanStack Router
  const rootRoute = createRootRoute({ component: () => ui });
  const router = createRouter({
    routeTree: rootRoute,
    history: createMemoryHistory(),
  });

  return render(
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>,
  );
};

const mockPost = {
  id: "1",
  userId: "user-123",
  username: "johndoe",
  displayname: "John Doe",
  imageUrl: "",
  color: "blue",
  likeCount: 10,
  isLiked: false,
  bookmarkCount: 5,
  isBookmarked: false,
  createdAt: "2024-01-01",
  updatedAt: "2024-01-01",
  word: "",
  definition: "",
  partOfSpeech: "noun",
  pronunciation: "",
  synonyms: [],
  antonyms: [],
  example: "",
  visibility: "public",
} satisfies Post;

describe("PostView Component", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders post content correctly", () => {
    renderWithProviders(<PostView post={mockPost} />);

    expect(screen.getByText("@johndoe")).toBeInTheDocument();
    expect(screen.getByText("John Doe")).toBeInTheDocument();
    expect(screen.getByText("10")).toBeInTheDocument(); // Like count
    expect(screen.getByText("5")).toBeInTheDocument(); // Bookmark count
  });

  it("handles optimistic like toggle and API call", async () => {
    vi.mocked(api.toggleLikePost).mockResolvedValueOnce();
    renderWithProviders(<PostView post={mockPost} />);

    const likeButton = screen.getByRole("button", { name: /10/i });

    // Click like
    fireEvent.click(likeButton);

    // Verify optimistic update (UI updates before API finishes)
    expect(screen.getByText("11")).toBeInTheDocument();

    // Verify API call
    expect(api.toggleLikePost).toHaveBeenCalledWith("1", false);

    await waitFor(() => {
      expect(likeButton).not.toBeDisabled();
    });
  });

  it("reverts like count if API fails", async () => {
    vi.mocked(api.toggleLikePost).mockRejectedValueOnce(
      new Error("Network Error"),
    );
    renderWithProviders(<PostView post={mockPost} />);

    const likeButton = screen.getByRole("button", { name: /10/i });

    fireEvent.click(likeButton);
    expect(screen.getByText("11")).toBeInTheDocument(); // Initially goes up

    // Wait for the catch block to trigger the revert
    await waitFor(() => {
      expect(screen.getByText("10")).toBeInTheDocument();
    });
  });

  it("opens the dialog when the expand button is clicked", async () => {
    renderWithProviders(<PostView post={mockPost} />);

    // The dialog content shouldn't be visible yet (depending on your Dialog implementation)
    // Radix Dialogs usually remove content from DOM or hide it.
    const expandButton = screen.getByRole("button", { name: /expand/i }); // Lucide icons usually need an aria-label or title

    fireEvent.click(expandButton);

    // Check for username inside the Dialog Title
    const dialogTitles = await screen.findAllByText("johndoe");
    expect(dialogTitles.length).toBeGreaterThan(0);
  });
});
