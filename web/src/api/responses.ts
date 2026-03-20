import type { Definition, Post, Profile } from "@/types";

export type APIError = "internal server error" | "";

export type Response<T> =
  | ({
      success: true;
    } & T)
  | {
      success: false;
      error: string;
    };

export type PostsResponse = {
  posts: Post[];
  currentPage: number;
  totalPages: number;
};

export type ProfileResponse = Profile & {
  postCount: number;
  followerCount: number;
  followingCount: number;
};

export type FollowResponse = {
  users: Profile[];
  currentPage: number;
  totalPages: number;
};

export type DefineResponse =
  | {
      type: "definitions";
      definitions: Definition[];
    }
  | {
      type: "suggestions";
      suggestions: string[];
    };
