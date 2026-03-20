import * as React from "react";
import { Separator as SeparatorPrimitive } from "radix-ui";

import { cn } from "@/lib/utils";

function Separator({
  children,
  className,
  orientation = "horizontal",
  decorative = true,
  ...props
}: React.ComponentProps<typeof SeparatorPrimitive.Root>) {
  return (
    <SeparatorPrimitive.Root
      data-slot="separator"
      decorative={decorative}
      orientation={orientation}
      className={cn(
        "relative data-vertical:self-stretch bg-border data-horizontal:w-full data-vertical:w-px data-horizontal:h-px shrink-0",
        className,
      )}
      {...props}
    >
      <div className="top-1/2 left-1/2 absolute bg-background -translate-1/2 -translate-y-1/2">
        {children}
      </div>
    </SeparatorPrimitive.Root>
  );
}

export { Separator };
