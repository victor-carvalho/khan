// +build ignore

// TEMPORARY AUTOGENERATED FILE: easyjson bootstapping code to launch
// the actual generator.

package main

import (
  "fmt"
  "os"

  "github.com/mailru/easyjson/gen"

  pkg "github.com/topfreegames/khan/api"
)

func main() {
  g := gen.NewGenerator("payload_easyjson.go")
  g.SetPkg("api", "github.com/topfreegames/khan/api")
  g.NoStdMarshalers()
  g.Add(pkg.EasyJSON_exporter_ApplyForMembershipPayload(nil))
  g.Add(pkg.EasyJSON_exporter_ApproveOrDenyMembershipInvitationPayload(nil))
  g.Add(pkg.EasyJSON_exporter_BasePayloadWithRequestorAndPlayerPublicIDs(nil))
  g.Add(pkg.EasyJSON_exporter_CreateClanPayload(nil))
  g.Add(pkg.EasyJSON_exporter_CreateGamePayload(nil))
  g.Add(pkg.EasyJSON_exporter_CreatePlayerPayload(nil))
  g.Add(pkg.EasyJSON_exporter_HookPayload(nil))
  g.Add(pkg.EasyJSON_exporter_InviteForMembershipPayload(nil))
  g.Add(pkg.EasyJSON_exporter_TransferClanOwnershipPayload(nil))
  g.Add(pkg.EasyJSON_exporter_UpdateClanPayload(nil))
  g.Add(pkg.EasyJSON_exporter_UpdateGamePayload(nil))
  g.Add(pkg.EasyJSON_exporter_UpdatePlayerPayload(nil))
  g.Add(pkg.EasyJSON_exporter_Validation(nil))
  if err := g.Run(os.Stdout); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
