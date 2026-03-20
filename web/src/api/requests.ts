import type { Definition, Visibility } from "@/types";

export type LoginRequest = {
  email: string;
  password: string;
};

export type RegisterRequest = {
  email: string;
  username: string;
  displayname: string;
  password: string;
};

export type RecentPostsRequest = {
  pageParam: number;
};

export type FollowingPostsRequest = {
  pageParam: number;
};

export type ProfilePostsRequest = {
  userId: string;
  pageParam: number;
};

export type CreatePostRequest = {
  content: Definition;
  visibility: Visibility;
};

export type FollowingRequest = {
  pageParam: number;
  userId: string;
};

export type FollowersRequest = {
  pageParam: number;
  userId: string;
};
