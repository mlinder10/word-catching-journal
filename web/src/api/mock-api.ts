/* eslint-disable @typescript-eslint/no-unused-vars */

import type { Definition, User } from "@/types";
import type {
  CreatePostRequest,
  FollowingPostsRequest,
  FollowingRequest,
  LoginRequest,
  ProfilePostsRequest,
  RecentPostsRequest,
  RegisterRequest,
} from "./requests";
import type {
  Response,
  PostsResponse,
  ProfileResponse,
  FollowResponse,
} from "./responses";

const USER_ID = crypto.randomUUID();
const REQUEST_DURATION = 1000;
const THROW_ERROR_PROBABILITY = 0;

export default class MockAPI {
  constructor() {}

  private async beforeEach(
    { dur = REQUEST_DURATION, prob = THROW_ERROR_PROBABILITY } = {
      dur: REQUEST_DURATION,
      prob: THROW_ERROR_PROBABILITY,
    },
  ) {
    await new Promise((res) => setTimeout(res, dur));
    if (Math.random() < prob) {
      throw new Error("mock error");
    }
  }

  // auth

  async validateToken(): Promise<Response<User>> {
    await this.beforeEach({ dur: 0 });
    return {
      success: true,
      id: USER_ID,
      email: "mattchristopherlinder@gmail.com",
      username: "mlinder",
      displayname: "Mittis",
      imageUrl: "https://picsum.photos/200/200",
      color: "#ff8326",
    };
  }

  async login(_: LoginRequest): Promise<Response<User>> {
    await this.beforeEach();
    return {
      success: true,
      id: USER_ID,
      email: "mattchristopherlinder@gmail.com",
      username: "mlinder",
      displayname: "Mittis",
      imageUrl: "https://picsum.photos/200/200",
      color: "#ff8326",
    };
  }

  async register(_: RegisterRequest): Promise<Response<User>> {
    await this.beforeEach();
    return {
      success: true,
      id: USER_ID,
      email: "mattchristopherlinder@gmail.com",
      username: "mlinder",
      displayname: "Mittis",
      imageUrl: "https://picsum.photos/200/200",
      color: "#ff8326",
    };
  }

  async logout(): Promise<void> {
    await this.beforeEach();
  }

  // feed

  async fetchRecentPosts(_: RecentPostsRequest): Promise<PostsResponse> {
    await this.beforeEach();
    return {
      posts: Array.from({ length: 10 }).map(() => ({
        id: crypto.randomUUID(),
        word: "Test",
        definition:
          "take measures to check the quality, performance, or reliability of (something), especially before putting it into widespread use or practice.",
        partOfSpeech: "verb",
        pronunciation: "test",
        synonyms: ["Try", "Attempt", "Experiment"],
        antonyms: [],
        example: "We must test our code before we deploy it.",
        visibility: "public",
        createdAt: new Date().toString(),
        updatedAt: new Date().toString(),
        likeCount: 23,
        bookmarkCount: 8,
        isLiked: true,
        isBookmarked: false,

        userId: USER_ID,
        username: "mlinder",
        displayname: "Mittis",
        imageUrl: "https://picsum.photos/200/200",
        color: "#ff8326",
      })),
      currentPage: 1,
      totalPages: 1,
    };
  }

  async fetchFollowingPosts(_: FollowingPostsRequest): Promise<PostsResponse> {
    await this.beforeEach();
    return {
      posts: Array.from({ length: 10 }).map(() => ({
        id: crypto.randomUUID(),
        word: "Test",
        definition:
          "take measures to check the quality, performance, or reliability of (something), especially before putting it into widespread use or practice.",
        partOfSpeech: "verb",
        pronunciation: "test",
        synonyms: ["Try", "Attempt", "Experiment"],
        antonyms: [],
        example: "We must test our code before we deploy it.",
        visibility: "public",
        createdAt: new Date().toString(),
        updatedAt: new Date().toString(),
        likeCount: 23,
        bookmarkCount: 8,
        isLiked: true,
        isBookmarked: false,

        userId: USER_ID,
        username: "mlinder",
        displayname: "Mittis",
        imageUrl: "https://picsum.photos/200/200",
        color: "#ff8326",
      })),
      currentPage: 1,
      totalPages: 1,
    };
  }

  // post action

  async toggleLikePost(_: string, __: boolean): Promise<void> {
    await this.beforeEach({ dur: 100, prob: 0 });
  }

  async toggleBookmarkPost(_: string, __: boolean): Promise<void> {
    await this.beforeEach({ dur: 100, prob: 0 });
  }

  // profile

  async fetchProfile(_: string): Promise<ProfileResponse> {
    await this.beforeEach({ dur: 1000 });
    return {
      id: crypto.randomUUID(),
      username: "mlinder",
      displayname: "Mittis",
      imageUrl: "https://picsum.photos/200/200",
      color: "#ff8326",
      postCount: 10,
      followerCount: 328,
      followingCount: 10012,
    };
  }

  async updateProfilePic(_: Uint8Array) {
    await this.beforeEach({ dur: 500 });
    // TODO:
    return "https://picsum.photos/200/200";
  }

  async fetchProfilePosts(_: ProfilePostsRequest): Promise<PostsResponse> {
    await this.beforeEach();
    return {
      posts: Array.from({ length: 10 }).map(() => ({
        id: crypto.randomUUID(),
        word: "Test",
        definition:
          "take measures to check the quality, performance, or reliability of (something), especially before putting it into widespread use or practice.",
        partOfSpeech: "verb",
        pronunciation: "test",
        synonyms: ["Try", "Attempt", "Experiment"],
        antonyms: [],
        example: "We must test our code before we deploy it.",
        visibility: "public",
        createdAt: new Date().toString(),
        updatedAt: new Date().toString(),
        likeCount: 23,
        bookmarkCount: 8,
        isLiked: true,
        isBookmarked: false,

        userId: USER_ID,
        username: "mlinder",
        displayname: "Mittis",
        imageUrl: "https://picsum.photos/200/200",
        color: "#ff8326",
      })),
      currentPage: 1,
      totalPages: 3,
    };
  }

  // follow

  async toggleFollowUser(_: string, __: boolean): Promise<void> {
    await this.beforeEach();
  }

  async fetchFollowing(_: FollowingRequest): Promise<FollowResponse> {
    await this.beforeEach();
    return {
      users: Array.from({ length: 20 }).map(() => ({
        id: crypto.randomUUID(),
        username: "mlinder",
        displayname: "Mittis",
        imageUrl: "https://picsum.photos/200/200",
        color: "#ff8326",
      })),
      currentPage: 1,
      totalPages: 1,
    };
  }

  async fetchFollowers(_: FollowingRequest): Promise<FollowResponse> {
    await this.beforeEach();
    return {
      users: Array.from({ length: 20 }).map(() => ({
        id: crypto.randomUUID(),
        username: "mlinder",
        displayname: "Mittis",
        imageUrl: "https://picsum.photos/200/200",
        color: "#ff8326",
      })),
      currentPage: 1,
      totalPages: 1,
    };
  }

  // define

  async defineWord(_: string): Promise<Definition[]> {
    await this.beforeEach();
    return Array.from({ length: 6 }).map(() => ({
      word: "Test",
      definition:
        "take measures to check the quality, performance, or reliability of (something), especially before putting it into widespread use or practice.",
      partOfSpeech: "verb",
      pronunciation: "test",
      synonyms: ["Try", "Attempt", "Experiment"],
      antonyms: [],
      example: "We must test our code before we deploy it.",
    }));
  }

  // create

  async createPost(_: CreatePostRequest): Promise<void> {
    await this.beforeEach();
  }
}
