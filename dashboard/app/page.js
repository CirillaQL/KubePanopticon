import NavBar from "@/components/NavBar";

export default function Home() {
  return (
      <div className="flex flex-col">
          <NavBar></NavBar>
          <div className="bg-[#e4e4e7] flex  justify-center">
              <div>
                  <p className="text-2xl">集群信息</p>
              </div>
          </div>
      </div>
  );
}
