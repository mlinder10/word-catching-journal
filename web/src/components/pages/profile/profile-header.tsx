import api from "@/api";
import ProfilePicture from "@/components/custom-ui/profile-picture";
import { useQuery } from "@tanstack/react-query";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { FollowersList, FollowingList } from "./follow-list";

type ProfileHeaderSectionProps = {
  userId: string;
};

export default function ProfileHeaderSection({
  userId,
}: ProfileHeaderSectionProps) {
  const { data, isPending, error } = useQuery({
    queryKey: ["user", userId],
    queryFn: async () => api.fetchProfile(userId),
  });

  if (isPending) return <ProfileHeaderSkeleton />;
  if (error || !data) return null;

  return (
    <section className="flex justify-between items-center px-4 py-4">
      <div className="flex gap-4">
        <ProfilePicture
          username={data.username}
          imageUrl={data.imageUrl}
          color={data.color}
        />
        <div>
          <p className="text-muted-foreground text-sm">@{data.username}</p>
          <p>{data.displayname}</p>
        </div>
      </div>

      <div className="flex gap-4 text-muted-foreground text-sm">
        <div className="flex flex-col items-center">
          <p>{data.postCount}</p>
          <p>Posts</p>
        </div>
        <Sheet>
          <SheetTrigger className="flex flex-col items-center">
            <p>{data.followingCount}</p>
            <p>Following</p>
          </SheetTrigger>
          <SheetContent>
            <SheetHeader>
              <SheetTitle>Following</SheetTitle>
            </SheetHeader>
            <FollowingList userId={userId} />
          </SheetContent>
        </Sheet>
        <Sheet>
          <SheetTrigger className="flex flex-col items-center">
            <p>{data.followerCount}</p>
            <p>Followers</p>
          </SheetTrigger>
          <SheetContent>
            <SheetHeader>
              <SheetTitle>Followers</SheetTitle>
            </SheetHeader>
            <FollowersList userId={userId} />
          </SheetContent>
        </Sheet>
      </div>
    </section>
  );
}

function ProfileHeaderSkeleton() {
  return (
    <section className="flex justify-between items-center px-4 py-4">
      <div className="flex gap-4">
        <Skeleton className="rounded-full w-16 aspect-square" />
        <div className="space-y-2">
          <Skeleton className="rounded-full w-20 h-4" />
          <Skeleton className="rounded-full w-32 h-4" />
        </div>
      </div>

      <div className="flex gap-4 text-muted-foreground text-sm">
        <div className="flex flex-col items-center gap-2">
          <Skeleton className="rounded-full w-8 h-4" />
          <Skeleton className="rounded-full w-16 h-4" />
        </div>
        <div className="flex flex-col items-center gap-2">
          <Skeleton className="rounded-full w-8 h-4" />
          <Skeleton className="rounded-full w-16 h-4" />
        </div>
        <div className="flex flex-col items-center gap-2">
          <Skeleton className="rounded-full w-8 h-4" />
          <Skeleton className="rounded-full w-16 h-4" />
        </div>
      </div>
    </section>
  );
}
