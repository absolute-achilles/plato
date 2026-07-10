"use client"

import { useState } from "react"
import { toast } from "sonner"

import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

interface DetailField {
  key: string
  label: string
  value: string
  type?: "text" | "email" | "tel" | "number" | "textarea"
}

interface DetailDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  title: string
  description?: string
  fields: DetailField[]
  onSave?: () => void
}

export function DetailDialog({
  open,
  onOpenChange,
  title,
  description,
  fields,
  onSave,
}: DetailDialogProps) {
  const [formData, setFormData] = useState(() =>
    Object.fromEntries(fields.map((f) => [f.key, f.value]))
  )
  const [isSaving, setIsSaving] = useState(false)

  const handleSave = async () => {
    setIsSaving(true)
    await new Promise((resolve) => setTimeout(resolve, 600))
    setIsSaving(false)
    toast.success("Changes saved (mock)", {
      description: "This is a demo action. Data will not persist.",
    })
    onSave?.()
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="clay-card max-w-lg border-0 sm:max-w-xl">
        <DialogHeader>
          <DialogTitle className="font-heading text-xl">{title}</DialogTitle>
          {description && <DialogDescription>{description}</DialogDescription>}
        </DialogHeader>

        <div className="grid gap-4 py-4">
          {fields.map((field) => (
            <div key={field.key} className="grid gap-2">
              <Label htmlFor={field.key} className="text-sm font-medium">
                {field.label}
              </Label>
              {field.type === "textarea" ? (
                <textarea
                  id={field.key}
                  value={formData[field.key]}
                  onChange={(e) =>
                    setFormData((prev) => ({
                      ...prev,
                      [field.key]: e.target.value,
                    }))
                  }
                  className="clay-input min-h-[80px] w-full resize-none p-3 text-sm"
                />
              ) : (
                <Input
                  id={field.key}
                  type={field.type || "text"}
                  value={formData[field.key]}
                  onChange={(e) =>
                    setFormData((prev) => ({
                      ...prev,
                      [field.key]: e.target.value,
                    }))
                  }
                  className="clay-input"
                />
              )}
            </div>
          ))}
        </div>

        <div className="flex justify-end gap-2">
          <Button
            variant="outline"
            onClick={() => onOpenChange(false)}
            className="clay-button rounded-xl border-border"
          >
            Close
          </Button>
          <Button
            onClick={handleSave}
            disabled={isSaving}
            className="clay-button bg-primary text-primary-foreground hover:bg-primary/90"
          >
            {isSaving ? "Saving..." : "Save Changes"}
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}
