package functions

import "fmt"

// This is an example. Replace thie following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	fmt.Fprintln(m.Out, `RUN apt-get update`)
	fmt.Fprintln(m.Out, `RUN apt-get install -y wget gnupg gnupg2 apt-transport-https`)
	fmt.Fprintln(m.Out, `RUN wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.asc.gpg`)
	fmt.Fprintln(m.Out, `RUN mv microsoft.asc.gpg /etc/apt/trusted.gpg.d/`)
	fmt.Fprintln(m.Out, `RUN wget -q https://packages.microsoft.com/config/debian/9/prod.list`)
	fmt.Fprintln(m.Out, `RUN mv prod.list /etc/apt/sources.list.d/microsoft-prod.list`)
	fmt.Fprintln(m.Out, `RUN chown root:root /etc/apt/trusted.gpg.d/microsoft.asc.gpg`)
	fmt.Fprintln(m.Out, `RUN chown root:root /etc/apt/sources.list.d/microsoft-prod.list`)
	fmt.Fprintln(m.Out, `RUN apt-get update`)
	fmt.Fprintln(m.Out, `RUN apt-get install -y azure-functions-core-tools`)

	return nil
}
