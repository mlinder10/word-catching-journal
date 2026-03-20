import type { User } from "@/types";
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
  DefineResponse,
} from "./responses";
import axios, { type AxiosInstance } from "axios";
import { parseNetworkError } from "@/lib/api";

// errors aren't handle here because they're caught by useQuery and useMutation
export default class API {
  private client: AxiosInstance;

  constructor(url: string, headers: Record<string, string> = {}) {
    this.client = axios.create({
      baseURL: url,
      headers,
      withCredentials: true,
    });
    // this.client.interceptors.response.use((res) => {
    //   console.log({
    //     path: res.config.url,
    //     data: res.data,
    //   });
    //   return res;
    // });
  }

  // auth

  async validateToken(): Promise<Response<User>> {
    try {
      const res = await this.client.get<User>("/refresh");
      return { success: true, ...res.data };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async login(data: LoginRequest): Promise<Response<User>> {
    try {
      const res = await this.client.post<User>("/login", data);
      return { success: true, ...res.data };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async register(data: RegisterRequest): Promise<Response<{ id: string }>> {
    try {
      const res = await this.client.post<string>("/register", data);
      return { success: true, id: res.data };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async logout(): Promise<Response<object>> {
    try {
      await this.client.post("/logout");
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async verifyEmail(code: string): Promise<Response<User>> {
    try {
      const res = await this.client.post<User>("/verify", { code });
      return { success: true, ...res.data };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async resendVerification(id: string): Promise<Response<object>> {
    try {
      await this.client.post(`/verify/resend/${id}`);
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async changePassword(
    oldPassword: string,
    newPassword: string,
  ): Promise<Response<object>> {
    try {
      await this.client.post(`/reset-password`, {
        oldPassword,
        newPassword,
      });
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async requestResetPassword(email: string): Promise<Response<object>> {
    try {
      await this.client.post(`/reset-password`, { email });
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  async resetPassword(
    code: string,
    newPassword: string,
  ): Promise<Response<object>> {
    try {
      await this.client.post(`/reset-password/${code}`, {
        password: newPassword,
      });
      return { success: true };
    } catch (error) {
      return {
        success: false,
        error: parseNetworkError(error),
      };
    }
  }

  // feed

  async fetchRecentPosts({
    pageParam,
  }: RecentPostsRequest): Promise<PostsResponse> {
    const res = await this.client.get<PostsResponse>(
      `/feed/recent?page=${pageParam}`,
    );
    return res.data;
  }

  async fetchFollowingPosts({
    pageParam,
  }: FollowingPostsRequest): Promise<PostsResponse> {
    const res = await this.client.get<PostsResponse>(
      `/feed/following?page=${pageParam}`,
    );
    return res.data;
  }

  // post action

  async toggleLikePost(postId: string, isLiked: boolean): Promise<void> {
    const res = await this.client.request({
      method: isLiked ? "DELETE" : "POST",
      url: `/post/${postId}/like`,
    });
    return res.data;
  }

  async toggleBookmarkPost(
    postId: string,
    isBookmarked: boolean,
  ): Promise<void> {
    const res = await this.client.request({
      method: isBookmarked ? "DELETE" : "POST",
      url: `/post/${postId}/bookmark`,
    });
    return res.data;
  }

  // profile

  async fetchProfile(userId: string): Promise<ProfileResponse> {
    const res = await this.client.get<ProfileResponse>(`/profile/${userId}`);
    return res.data;
  }

  // TODO
  async updateProfilePic() {}

  async fetchProfilePosts({
    userId,
    pageParam,
  }: ProfilePostsRequest): Promise<PostsResponse> {
    const res = await this.client.get<PostsResponse>(
      `/profile/${userId}/posts?page=${pageParam}`,
    );
    return res.data;
  }

  // follow

  async toggleFollowUser(userId: string, isFollowing: boolean): Promise<void> {
    const res = await this.client.request({
      method: isFollowing ? "DELETE" : "POST",
      url: `/profile/${userId}/follow`,
    });
    return res.data;
  }

  async fetchFollowing({
    userId,
    pageParam,
  }: FollowingRequest): Promise<FollowResponse> {
    const res = await this.client.get<FollowResponse>(
      `/profile/${userId}/following?page=${pageParam}`,
    );
    return res.data;
  }

  async fetchFollowers({
    userId,
    pageParam,
  }: FollowingRequest): Promise<FollowResponse> {
    const res = await this.client.get<FollowResponse>(
      `/profile/${userId}/followers?page=${pageParam}`,
    );
    return res.data;
  }

  // define

  async defineWord(word: string): Promise<DefineResponse> {
    const res = await this.client.get<DefineResponse>(`/define/${word}`);
    return res.data;
  }

  // create post

  async createPost(data: CreatePostRequest): Promise<void> {
    await this.client.post(`/post`, data);
  }

  async deletePost(postId: string): Promise<void> {
    await this.client.delete(`/post/${postId}`);
  }
}
